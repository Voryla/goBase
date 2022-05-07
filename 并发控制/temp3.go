package main

import (
	"fmt"
	"sync"
)

func main() {
	for {
		var x, y = 0, 0
		c := sync.WaitGroup{}
		c.Add(2)
		go func() {
			x = 1
			fmt.Printf("%d ", y)
			c.Done()
		}()

		go func() {
			y = 1
			fmt.Printf("%d ", x)
			c.Done()
		}()
		c.Wait()
		fmt.Println()
	}
}
