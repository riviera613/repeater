package repeater

import (
	"github.com/panjf2000/ants/v2"
	"log"
	"math"
	"sort"
	"time"
)

// TestCase Single test case
type TestCase struct {
	Name         string
	Func         func() error
	Concurrence  int64
	TotalCount   int64
	SuccessCount int64
	Total        float64
	Avg          float64
	P95          float64
	P99          float64
	IsFinish     bool
}

// Init init
func (c *TestCase) Init(inputFunc *InputFunc, inputParam *InputParam) {
	c.Name = inputFunc.Name
	c.Func = inputFunc.Func
	c.Concurrence = inputParam.Concurrence
	c.TotalCount = inputParam.TotalCount
	c.IsFinish = false
}

// Process run stress test, and save result
func (c *TestCase) Process() {
	defer ants.Release()
	c.IsFinish = true
	log.Printf("[%s] Start", c.Name)
	ch := make(chan float64, c.TotalCount)
	res := make([]float64, c.TotalCount)
	pool, _ := ants.NewPool(int(c.Concurrence))
	start := time.Now().UnixNano()
	for i := 0; i < int(c.TotalCount); i++ {
		if err := pool.Submit(c.testWrapper(ch)); err != nil {
			log.Printf("[%s] Submit task error: %s", c.Name, err.Error())
		}
	}
	for i := 0; i < int(c.TotalCount); i++ {
		res[i] = <-ch
	}
	end := time.Now().UnixNano()
	c.Total = float64(end-start) / 1e9
	c.stats(res)
	c.IsFinish = true
}

func (c *TestCase) testWrapper(ch chan float64) func() {
	return func() {
		duration := float64(-1)
		defer func() {
			if r := recover(); r != nil {
				log.Printf("[%s] Test recover from panic", c.Name)
			}
			ch <- duration
		}()

		start := time.Now().UnixNano()
		err := c.Func()
		end := time.Now().UnixNano()
		if err == nil {
			duration = float64(end-start) / 1e9
		} else {
			log.Printf("[%s] Test error: %s", c.Name, err.Error())
		}
	}
}

func (c *TestCase) stats(res []float64) {
	sort.Float64s(res)
	for _, v := range res {
		if v >= 0 {
			c.Avg += v
			c.SuccessCount += 1
		}
	}
	if c.SuccessCount > 0 {
		c.Avg /= float64(c.SuccessCount)
		c.P95 = res[int64(math.Floor(float64(c.SuccessCount)*95/100))]
		c.P99 = res[int64(math.Floor(float64(c.SuccessCount)*99/100))]
	} else {
		c.Avg = -1
		c.P95 = -1
		c.P99 = -1
	}
}
