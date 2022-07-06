package aboutArray

func search1(nums []int, target int) int {
	right := len(nums) - 1
	left := 0
	answer := 0
	// 1.通过二分查找，找到目标数所在的下标，随后向左向右遍历，寻找相同的数的个数
	for left <= right {
		mid := (right-left)/2 + left
		if nums[mid] == target {
			// 向右寻找相同的数字
			for i := mid + 1; i < len(nums); i++ {
				if nums[i] == target {
					answer++
				} else {
					break
				}
			}
			// 向左寻找相同的数字
			for i := mid; i >= 0; i-- {
				if nums[i] == target {
					answer++
				} else {
					break
				}
			}
			return answer
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return answer
}
