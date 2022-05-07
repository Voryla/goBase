package main

import (
	"fmt"
	"time"
)

func main() {
	a := 1 // 因为在闭包函数后对a依旧有操作，所以此时a变成了引用值
	b := 2
	go func() {
		fmt.Println(a, b)
	}()
	a = 99
	time.Sleep(time.Second)
}

func re() func() {
	w := 1
	return func() {
		fmt.Println(w)
	}
}
