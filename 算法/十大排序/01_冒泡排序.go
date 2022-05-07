package main

import "fmt"

func main() {
	defer fmt.Println("a")
	defer fmt.Println("b")
	defer fmt.Println("c")
}

func sortArray(nums []int) []int {
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := 0; j < length-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}
