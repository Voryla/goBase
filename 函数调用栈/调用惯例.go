package main

import "fmt"

func myFunction(a, b int) (int, int) {
	return a + b, a - b
}

func main() {
	var a = map[int]int{}
	a[1] = 1
	fmt.Println(a)
	var c = map[string]byte{}
	c["1"] = 2
	m := make(map[string]byte)
	fmt.Println(m)
}
