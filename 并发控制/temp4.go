package main

import (
	"fmt"
)

func f(n int) (r int) {
	defer func() {
		r += n
		fmt.Println("r")

	}()

	var f func()
	f = func() {
		fmt.Println("f")
		r += 2
	}
	defer f()

	return n + 1
}

func main() {
	fmt.Println(f(3))
}
