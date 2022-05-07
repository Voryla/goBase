package main

import (
	"fmt"
)

func add(a, b int) int
func swap(a, b int) (ret0, ret1 int)
func main() {
	fmt.Println(add(10, 11))
	a, b := 1, 2
	fmt.Printf("Befor Swap:a:%d,b:%d\n", a, b)
	a, b = swap(a, b)
	fmt.Printf("After Swap:a:%d,b:%d\n", a, b)
}
