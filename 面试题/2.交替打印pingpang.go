package main

import (
	"fmt"
	"time"
)

func main() {
	ping, pang := make(chan string), make(chan string)
	go func() {
		for {
			pang <- "pang"
			fmt.Println(<-ping)
			time.Sleep(1 * time.Second)
		}
	}()
	go func() {
		for {
			time.Sleep(2 * time.Second)
			fmt.Println(<-pang)
			ping <- "ping"
		}
	}()
	time.Sleep(30 * time.Second)
}
