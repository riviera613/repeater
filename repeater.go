package repeater

import (
	"fmt"
	"github.com/jedib0t/go-pretty/table"
	"log"
	"os"
)

type Repeater struct {
	TestCases []*TestCase
}

// Init init
func (r *Repeater) Init(inputFuncList []*InputFunc, inputParamList []*InputParam) {
	r.TestCases = make([]*TestCase, 0)
	for _, inputFunc := range inputFuncList {
		if inputFunc.Name == "" || inputFunc.Func == nil {
			log.Printf("Invalid input func: %v", inputFunc)
			continue
		}
		for _, inputParam := range inputParamList {
			if inputParam.TotalCount <= 0 || inputParam.Concurrence <= 0 {
				log.Printf("Invalid input params: %v", inputParam)
				continue
			}
			testCase := &TestCase{}
			testCase.Init(inputFunc, inputParam)
			r.TestCases = append(r.TestCases, testCase)
		}
	}
}

// Process run all test cases
func (r *Repeater) Process() {
	for _, testCase := range r.TestCases {
		testCase.Process()
	}
}

// ToCsv output result to csv file
func (r *Repeater) ToCsv(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	_, _ = file.WriteString("Name,Concurrence,Test Count,Success Count,Success Rate,Total,Avg,P95,P99\n")
	for _, c := range r.TestCases {
		csvRow := fmt.Sprintf("%s,%d,%d,%d,%.2f,%f,%f,%f,%f\n", c.Name, c.Concurrence, c.TotalCount, c.SuccessCount, float64(c.SuccessCount)/float64(c.TotalCount), c.Total, c.Avg, c.P95, c.P99)
		if _, err = file.WriteString(csvRow); err != nil {
			return err
		}
	}
	return nil
}

// Render render result to a string
func (r *Repeater) Render() string {
	t := table.Table{}
	header := table.Row{"Name", "Concurrence", "Total Count", "Success Count", "Success Rate", "Total", "Avg", "P95", "P99"}
	t.AppendHeader(header)
	for _, c := range r.TestCases {
		row := table.Row{c.Name, c.Concurrence, c.TotalCount, c.SuccessCount, float64(c.SuccessCount) / float64(c.TotalCount), c.Total, c.Avg, c.P95, c.P99}
		t.AppendRow(row)
	}
	return t.Render()
}

// NewRepeater get a new object
func NewRepeater(inputFuncList []*InputFunc, inputParamList []*InputParam) *Repeater {
	r := &Repeater{}
	r.Init(inputFuncList, inputParamList)
	return r
}
