package main

import "fmt"

type temp struct {
	x int
	y *int
}

// 递增运算符++和递减运算符--的优先级低于解引用运算符*和取地址运算符&，解引用运算符和取地址运算符的优先级低于选择器.中的属性选择操作符。
// 选择器.>解引用运算符*>递增++，递减--运算符
func main() {
	var t temp
	p := &t.x             // <=> p := &(t.x)
	fmt.Printf("%T\n", p) // *int

	*p++ // <=> (*p)++
	*p-- // <=> (*p)--

	t.y = p
	a := *t.y             // <=> *(t.y)
	fmt.Printf("%T\n", a) // int
}
