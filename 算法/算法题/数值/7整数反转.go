package 数值

import (
	"math"
	"strconv"
)

func re(x int) (rev int) {
	for x != 0 {
		if rev < math.MinInt32/10 || rev > math.MaxInt32/10 {
			return 0
		}
		digit := x % 10
		x /= 10
		rev = rev*10 + digit
	}
	return

}

func reverse(x int) int {
	isN := false
	if x < 0 {
		isN = true
		x *= -1
	}
	s := strconv.Itoa(x)
	temp := ""
	for i := len(s) - 1; i >= 0; i-- {
		temp += string(s[i])
	}
	x, _ = strconv.Atoi(temp)
	if isN {
		x *= -1
	}
	if x > math.MaxInt32 || x < math.MinInt32 {
		return 0
	}
	return x
}
