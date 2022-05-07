package main

import (
	"fmt"
)

func main() {
	a := 1
	b := 1.1
	//c := a + b  // a为整数，而b为浮点数 mismatched types int and float64

	q := 1
	//c := q + 1.1// q被推断为整数，而1.1是浮点数，q不能被强制转换为浮点数，相加失效 constant 1.1 truncated to integer。
	fmt.Println(a, b, q)
	//c := 1 / 3
	//var df int = -5
	//var le = time.Now().Add(time.Duration(df))
}
