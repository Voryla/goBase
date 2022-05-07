package main

import (
	"fmt"
	"time"
)

func main() {
	var a int = 0
	go func() {
		for {
			a = a + 1
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for {
			print1(a)
		}
	}()
	select {}
}
func print1(a int) {
	fmt.Println(a)
}
