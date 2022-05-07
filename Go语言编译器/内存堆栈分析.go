package main

import (
	"fmt"
)

func main() {
	x, y := 0.3, 0.6
	z := x + y
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	c := 1
	go stack(&c)
	c += 3
	fmt.Println(c)
}

func stack(i *int) {
	*i++
}

func do() {
	a := 1
	func() {
		a = 2
	}()
	a = 3
}
