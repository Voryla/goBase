package aboutArray

func moveZeroes(nums []int) {
	count := 0
	for i := 0; i < len(nums)-count; i++ {
		if nums[i] != 0 {
			continue
		}
		for j := count + 1; j < len(nums)-count; j++ {
			if nums[j] != 0 {
				nums[j], nums[i] = nums[i], nums[j]
				count++
			}
		}
	}
}
