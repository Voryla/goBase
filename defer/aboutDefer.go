package main

import (
	"fmt"
)

func main() {
	moveZeroes([]int{5, 1, 0, 3, 12})
}
func A1(a int) int {
	fmt.Println(a)
	return 3
}
func moveZeroes(nums []int) {
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}
