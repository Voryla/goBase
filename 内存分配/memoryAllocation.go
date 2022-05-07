package main

import "fmt"

func main() {
	test()
	var f = test
	f()
	f1 := test
	f1()
}

func test() {
	fmt.Println("a")
}
