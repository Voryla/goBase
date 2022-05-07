package main

import (
	"fmt"
	"os"
	"time"
)

// 我们可以通过调用os.Exit函数从任何函数里退出一个程序。 os.Exit函数调用接受一个int代码值做为参数并将此代码返回给操作系统。
func main() {
	defer fmt.Println("exit") // 该defer不会得到执行
	go func() {
		time.Sleep(time.Second)
		os.Exit(1) // 直接退出程序，不会执行之前压栈的defer
	}()
	select {}
}
