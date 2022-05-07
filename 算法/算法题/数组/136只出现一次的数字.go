package aboutArray

func singleNumber(nums []int) int {
	val := 0
	for _, num := range nums {
		// 异或运算相同的位为减，不同的位为加
		val ^= num
	}
	return val
}
