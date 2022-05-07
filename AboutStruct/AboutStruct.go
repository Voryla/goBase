package main

import "fmt"

type myStruct struct {
	a string
	b int
	c byte
}

func main() {
	var s = struct {
		a string
		b int
		c byte
	}{
		a: "hello",
		b: 3,
		c: 1,
	}
	fmt.Println(s)

	var my myStruct
	my = s
	fmt.Println(my)
}
