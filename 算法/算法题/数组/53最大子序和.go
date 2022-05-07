package aboutArray

import "C"
import (
	"fmt"
)

func main() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3, 4)
	fmt.Println(s)
}
func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
