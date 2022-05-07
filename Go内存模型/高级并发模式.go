package main

import (
	"fmt"
	"time"
)

func main() {
	//advancedGoConcurrencyPatterns()
	//advancedGoConcurrencyPatterns1()
	//advancedGoConcurrencyPatterns2()
	f := fib()
	// Function calls are evaluated left-to-right.
	fmt.Println(f(), f(), f(), f(), f())
}

func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

// Go 高级并发模式

type Ball struct {
	hits int
}

func advancedGoConcurrencyPatterns() {
	table := make(chan *Ball)
	go player("ping", table)
	go player("pong", table)
	table <- new(Ball)
	time.Sleep(time.Second)
	<-table
}

func advancedGoConcurrencyPatterns1() {
	table := make(chan *Ball)
	go player("ping", table)
	go player("pong", table)
	//table <- new(Ball)
	time.Sleep(time.Second)
	<-table
}

func advancedGoConcurrencyPatterns2() {
	table := make(chan *Ball)
	go player("ping", table)
	go player("pong", table)
	table <- new(Ball)
	//time.Sleep(time.Second)
	<-table
	panic("show me the stacks")
}

func player(name string, table chan *Ball) {
	for {
		ball := <-table
		ball.hits++
		fmt.Println(name, ball.hits)
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}
