package aboutArray

import (
	"fmt"
)

func tmain() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}
func merge(nums1 []int, m int, nums2 []int, n int) {
	temp := make([]int, 0)
	indexNums1, indexNums2 := 0, 0
	for indexNums1 < m && indexNums2 < n {
		if nums1[indexNums1] < nums2[indexNums2] {
			temp = append(temp, nums1[indexNums1])
			indexNums1++
		} else {
			temp = append(temp, nums2[indexNums2])
			indexNums2++
		}
	}
	if indexNums1 < m {
		temp = append(temp, nums1[indexNums1:]...)
	}
	if indexNums2 < n {
		temp = append(temp, nums2[indexNums2:]...)
	}
	copy(nums1, temp)
}
