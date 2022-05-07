package main

import (
	"fmt"
)

func main() {
	var a *byte
	fmt.Printf("%p\n", &a)
	for i := 0; i < 1000; i++ {
		c := int64(3)
		fmt.Printf("%p\n", &c)
		fmt.Println()
	}
}
