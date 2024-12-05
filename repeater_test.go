package repeater

import (
	"errors"
	"math/rand"
	"testing"
	"time"
)

func forTest() error {
	rand.NewSource(time.Now().UnixNano())
	n := time.Duration(rand.Intn(128))
	if n > 64 {
		return errors.New("timeout")
	}
	time.Sleep(n * time.Millisecond)
	return nil
}

func forTestPanic() error {
	panic("test panic")
}

func TestRepeater_Process(t *testing.T) {
	_repeater := NewRepeater([]*InputFunc{
		{Name: "forTest", Func: forTest},
		{Name: "forTestPanic", Func: forTestPanic},
		{Name: "", Func: nil},
	}, []*InputParam{
		{Concurrence: 5, TotalCount: 10},
		{Concurrence: -1, TotalCount: -1},
	})
	_repeater.Process()
	_repeater.Render()
	_ = _repeater.ToCsv("test.csv")
}
