package main

import (
	"fmt"
	"runtime"
	"time"
)

//%d表示输出十进制数字，+表示输出数值的符号。这里不表示取反。
func main() {
	//i := -5
	//j := +5
	//fmt.Printf("%+d %+d", i, j)
	var x int
	threads := runtime.GOMAXPROCS(0)
	fmt.Println("threads=", threads)
	for i := 0; i < threads; i++ {
		go func() {
			for {
				x++
				fmt.Println("x=", x)

			}
		}()
	}
	time.Sleep(time.Second)
	fmt.Println("x=", x)

}
