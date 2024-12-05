# Repeater
A simple stress test framework for Golang

English | [中文](README_ZH.md)

## Install
```powershell
go get -u riviera613/reapter
```

## How to use
```go
package main

import (
	"errors"
	"fmt"
	"github.com/riviera613/repeater"
	"math/rand"
	"time"
)

// forTest a function for test, generate a random integer N between 0 and 1024
// if N > 1000, return error; otherwise, sleep N ms
func forTest() error {
	n := time.Duration(rand.Intn(1024))
	if n > 1000 {
		return errors.New("timeout")
	}
	time.Sleep(n * time.Millisecond)
	return nil
}

func main() {
	_repeater := repeater.NewRepeater([]*repeater.InputFunc{
		{Name: "forTest", Func: forTest},
	}, []*repeater.InputParam{
		{Concurrence: 1, TotalCount: 5},
		{Concurrence: 10, TotalCount: 50},
		{Concurrence: 100, TotalCount: 500},
	})
	_repeater.Process()
	fmt.Println(_repeater.Render())   // output result to terminal
	_ = _repeater.ToCsv("result.csv") // output result to csv
}
```

## Output
```powershell
+---------+-------------+-------------+---------------+--------------+----------+--------------------+----------+----------+
| NAME    | CONCURRENCE | TOTAL COUNT | SUCCESS COUNT | SUCCESS RATE |    TOTAL |                AVG |      P95 |      P99 |
+---------+-------------+-------------+---------------+--------------+----------+--------------------+----------+----------+
| forTest |           1 |           5 |             5 |            1 | 2.973226 | 0.5945933999999999 | 0.931145 | 0.931145 |
| forTest |          10 |          50 |            50 |            1 | 3.390408 |         0.58240476 | 0.967909 | 0.983813 |
| forTest |         100 |         500 |           488 |        0.976 | 2.990706 | 0.4858624651639343 | 0.926045 | 0.975087 |
+---------+-------------+-------------+---------------+--------------+----------+--------------------+----------+----------+
```