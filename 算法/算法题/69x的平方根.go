package main

import (
	"fmt"
	"math"
)

func main() {
	s := int64(math.Exp(0.5 * math.Log(float64(4))))
	fmt.Println(s)
}

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	ans := int(math.Exp(0.5 * math.Log(float64(x))))
	// 判断返回 ans+1 还是 ans
	if (ans+1)*(ans+1) <= x {
		return ans + 1
	}
	return ans
}
