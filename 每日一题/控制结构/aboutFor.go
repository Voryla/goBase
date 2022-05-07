package main

import "fmt"

func main() {
	var a = [5]int{1, 2, 3, 4, 5}
	var r [5]int

	for i, v := range a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}
	for i, i2 := range a {
		go func() {
			fmt.Println(i2)
			fmt.Println(i)
		}()

	}
	fmt.Println("r = ", r)
	fmt.Println("a = ", a)
}
