package main

import (
	"fmt"
	"runtime"
)

// 我们可以通过调用runtime.Goexit函数退出一个goroutine。 runtime.Goexit函数没有参数。
func main() {
	c := make(chan int)
	go func() {
		defer func() { c <- 1 }()
		defer fmt.Println("Go")
		func() {
			defer fmt.Println("C")
			runtime.Goexit() // 立即退出Go程，并执行当前Go程压栈的defer
		}()
		fmt.Println("Java") // 该语句无法被执行
	}()
	<-c
}
