package main

import (
	"testing"
)

/*
一下代码可以对函数进行测试
go test .\内联函数_test.go -bench=BenchmarkMax

当某个函数前添加该注释：
代表当前函数是禁止进行函数内联优化的
//go:noinline
*/
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

var Result int

func BenchmarkMax(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = max(-1, i)
	}
	Result = r
}
