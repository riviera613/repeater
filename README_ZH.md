# Repeater
一个简单的Golang语言压力测试框架

[English](README.md) | 中文

## 安装
```powershell
go get -u riviera613/reapter
```

## 如何使用
```go
package main

import (
	"errors"
	"fmt"
	"github.com/riviera613/repeater"
	"math/rand"
	"time"
)

// ForTest 一个用于测试的方法，在0-1024之间生成随机数
// 如果结果大于1000，就认为是报错；否则就延迟对应的毫秒数
func ForTest() error {
	n := time.Duration(rand.Intn(1024))
	if n > 1000 {
		//
		return errors.New("timeout")
	}
	time.Sleep(n * time.Millisecond)
	return nil
}

func main() {
	_repeater := repeater.NewRepeater([]*repeater.InputFunc{
		{Name: "ForTest", Func: ForTest},
	}, []*repeater.InputParam{
		{Concurrence: 1, TotalCount: 5},
		{Concurrence: 10, TotalCount: 50},
		{Concurrence: 100, TotalCount: 500},
	})
	_repeater.Process()
	fmt.Println(_repeater.Render())   // 输出结果到控制台
	_ = _repeater.ToCsv("result.csv") // 输出结果到指定csv文件
}
```

## 输出示例
```powershell
+---------+-------------+-------------+---------------+--------------+----------+--------------------+----------+----------+
| NAME    | CONCURRENCE | TOTAL COUNT | SUCCESS COUNT | SUCCESS RATE |    TOTAL |                AVG |      P95 |      P99 |
+---------+-------------+-------------+---------------+--------------+----------+--------------------+----------+----------+
| ForTest |           1 |           5 |             5 |            1 | 2.973226 | 0.5945933999999999 | 0.931145 | 0.931145 |
| ForTest |          10 |          50 |            50 |            1 | 3.390408 |         0.58240476 | 0.967909 | 0.983813 |
| ForTest |         100 |         500 |           488 |        0.976 | 2.990706 | 0.4858624651639343 | 0.926045 | 0.975087 |
+---------+-------------+-------------+---------------+--------------+----------+--------------------+----------+----------+
```
