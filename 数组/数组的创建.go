package main

import (
	"fmt"
	"unsafe"
)

func main() {

	a := [3]int{1, 2, 3}
	b := 8
	_ = a[b] // 在编译时被检测到数组越界
	var c int64 = 3
	fmt.Printf("%d", unsafe.Sizeof(c))

}

// 创建数组的三种方式
func f1() {
	var arr [3]int
	var arr2 = [4]int{1, 2, 3, 4}
	arr3 := [...]int{1, 2, 3}
	fmt.Println(arr, arr2, arr3)
}
