package main

import "fmt"

type MyInt int

func (v MyInt) Twice() int

func main() {
	c := MyInt(3)
	v := c.Twice()
	fmt.Println(v)
}
