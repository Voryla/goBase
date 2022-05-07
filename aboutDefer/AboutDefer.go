package main

import (
	"fmt"
)

func main() {
	//fmt.Println(deferFunc())
	//executePanic()
	//fmt.Println("Main block is executed completely")
	//println(deferFunc4())
	//println(deferFunc())
	//executePanicWithRecover()
	//fmt.Println("Main block is executed completely")
	//deferIsLIFO()

	i := f()
	//i := f1()
	fmt.Printf("main: i=%d, g=%d\n", i, g)
}

// 在executePanic函数中，手动执行panic函数触发了异常。
//当异常触发后，函数仍然会调用defer中的函数，然后异常退出。
//输出如下，表明调用了defer中的函数，并且main函数将不能正常运 行，程序异常退出打印出栈追踪信息。
func executePanic() {
	defer fmt.Println("defer func")
	panic("This is Panic Situation")
	fmt.Println("The function executes Completely")
}

// 当在defer函数中使用recover进行异常捕获后，程序 将不会异常退出，并且能够执行正常的main函数流程
func executePanicWithRecover() {
	defer func() {
		if errMsg := recover(); errMsg != nil {
			fmt.Println(errMsg)
		}
		fmt.Println("This is recovery function")
	}()
	panic("This is Panic Situation")
	fmt.Println("The function executes Completely")
}
func deferFunc4() (t int) {
	defer func(i int) {
		i = t
	}(t)
	return 2
}

func deferFunc() (x int) {
	defer func() {
		x++
	}()
	return 2
}

func deferIsLIFO() {
	for i := 0; i < 5; i++ {
		defer fmt.Println("start", i)
	}
}

var g = 100

func f() (r int) {
	defer func() {
		g = 200
	}()
	fmt.Printf("f:g=%d\n", g)
	return g
}

func f1() (r int) {
	r = g
	defer func() {
		r = 200
	}()
	r = 0
	return r
}
