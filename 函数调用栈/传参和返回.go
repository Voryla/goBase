package main

import "fmt"

func incr(a int) int {
	var b int
	defer func() {
		a++
		b++
	}()
	a++
	b = a
	return b
}
func main() {
	var a, b int
	b = incr(a)
	fmt.Println(a, b)
	var z, w int
	w = incr1(z)
	fmt.Println(z, w)
}

func incr1(a int) (b int) {
	defer func() {
		a++
		b++
	}()
	a++
	return a
}
