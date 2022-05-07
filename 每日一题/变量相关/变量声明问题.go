package main

import "fmt"

func main() {
	// 下面代码中，x 已声明，y 没有声明，判断每条语句的对错。
	// 变量的声明。1.错，x 已经声明，不能使用 :=；2.对；3.对，当多值赋值时，:= 左边的变量无论声明与否都可以；4.错，y 没有声明。
	//x := 3
	//x, _ := finit()
	//x, _ = finit()
	//x, y := finit()
	//x, y = finit()
	//_ = y
	var a int = 32.0
	fmt.Println(a)
	var b int8 = 32.0
	fmt.Println(b)
	var c int16 = 3_2.
	fmt.Println(c)
	var d int32 = 32.0
	fmt.Println(d)
	var e int64 = 32.0
	fmt.Println(e)
	var f float64 = .03
	fmt.Println(f)
	s := new([]int)
	*s = append(*s, 1)

}
func finit() (int, int) {
	return 1, 2
}

type s struct {
	data []int
}
