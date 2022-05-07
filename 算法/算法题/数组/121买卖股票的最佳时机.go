package aboutArray

import "math"

func main() {
	println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

func maxProfit(prices []int) int {
	min := math.MaxInt
	max := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else if prices[i]-min > max {
			max = prices[i] - min
		}
	}
	return max
}
