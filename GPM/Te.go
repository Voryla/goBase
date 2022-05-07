package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {
	go func() {
		time.Sleep(4 * time.Second)
		fmt.Println("over")
	}()
	time.Sleep(1 * time.Second)
	c := make(chan struct{})
	fmt.Println(c)
	ints := []int{2, 53, 3, 5}
	sort.Ints(ints)
	for i := 0; i < len(ints); i++ {
		fmt.Println(ints[i])
	}
}
