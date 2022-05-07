package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var x int
	var b int
	threads := runtime.GOMAXPROCS(0)
	fmt.Println("threads=", threads)
	for i := 0; i < threads; i++ {
		go func() {
			for {
				x++
				b = x
			}
		}()
	}
	time.Sleep(time.Second)
	fmt.Println("x=", x)
	fmt.Println("x=", b)
}
