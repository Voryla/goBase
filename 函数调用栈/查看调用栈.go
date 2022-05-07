package main

import (
	"fmt"
	"runtime/debug"
)

type user struct {
	name string
	age  int
}

func main() {

	//a(99)
	u := create()
	s := u.name
	fmt.Println(s)
}

func create() user {
	u := &user{"zw", 1}
	return *u
}
func a(num int) {
	debug.PrintStack()
}
