package main

import "fmt"

// 下面代码下划线处可以填入哪个选项？
func main() {
	var s1 []int
	var s2 = []int{}
	if s1 == nil {
		fmt.Println("yes nil")
	} else {
		fmt.Println("no nil")
	}
	if s2 == nil {
		fmt.Println("yes nil")
	} else {
		fmt.Println("no nil")
	}
}

//``
//
//- A. s1
//- B. s2
//- C. s1、s2 都可以
