package main

import "fmt"

func main() {
	var b byte = '\177'
	fmt.Println(b)
	fmt.Println('\017')
	slice1 := []int{1, 2}
	testSlice(slice1)
	fmt.Println(slice1)
}

func testSlice(slice []int) {
	slice[0] = 2
	slice[1] = 1
}
