package main

import "fmt"

func main() {
	s := [3]int{1, 2, 3}
	a := s[:0]
	fmt.Printf("%d %d\n", len(a), cap(a))
	
	// 在操作符 [i:j] 中，如果 i 省略，默认 0，如果 j 省略，默认底层数组的长度，截取得到的切片长度和容量计算方法是 j-i、l-i
	b := s[:2]
	fmt.Printf("%d %d\n", len(b), cap(b))

	// 操作符 [i:j:k]，k 主要是用来限制切片的容量，但是不能大于数组的长度 l，截取得到的切片长度和容量计算方法是 j-i、k-i
	c := s[1:2:cap(s)]
	fmt.Printf("%d %d\n", len(c), cap(c))

}
