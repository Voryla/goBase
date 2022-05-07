package aboutArray

func majorityElement(nums []int) int {
	temp := make(map[int]int)
	max, val := 0, 0
	for i := 0; i < len(nums); i++ {
		temp[nums[i]]++
		data := temp[nums[i]]
		if data > max {
			max = data
			val = nums[i]
		}
	}
	return val
}
