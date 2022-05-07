package main

import "fmt"

func main() {
	//var slice1 []int
	//var slice2 []int = make([]int, 5)
	//var slice2 []int = make([]int, 5, 7)
	numbers := []int{1: 3, 2, 3, 4: 9, 5, 6}
	fmt.Println(numbers)
	numbers = numbers[3:4:5]
	fmt.Println(numbers)
}
