package main

import "fmt"

func main() {
	a := 0
	f := func() {}
	f = func() {
		b := 1
		fmt.Printf("a: %p----b: %p\n", &a, &b)
		f()
	}
	go f()
	select {}
}
