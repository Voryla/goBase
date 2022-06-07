package main

import "fmt"

func main() {
	/*
		- 零切片
		我们把切片内部数组的元素都是零值或者底层数组的内容就全是 `nil`的切片叫做零切片，使用`make`创建的、长度、容量都不为0的切片就是零值切片：
	*/
	slice := make([]int, 5)   // 0 0 0 0 0
	slice1 := make([]*int, 5) // nil nil nil nil nil
	fmt.Println(slice, slice1)
	/*
		- `nil`切片
		`nil`切片的长度和容量都为`0`，并且和`nil`比较的结果为`true`，采用直接创建切片的方式、`new`创建切片的方式都可以创建`nil`切片
	*/
	var slice2 []int
	fmt.Println(nil == slice2)
	var slice3 = *new([]int)
	fmt.Println(nil == slice3)
	fmt.Println(slice2, slice3)
	slice2 = append(slice2, 2, 2)
	slice3 = append(slice3, 3, 3)
	fmt.Println(slice2)
	fmt.Println(slice3)
	/*
		- 空切片
		空切片的长度和容量也都为`0`，但是和`nil`的比较结果为`false`，因为所有的空切片的数据指针都指向同一个地址 `0xc42003bda0`；使用字面量、`make`可以创建空切片：
	*/
	var slice4 = []int{}
	fmt.Println(nil == slice4)
	var slice5 = make([]int, 0)
	fmt.Println(nil == slice5)
	fmt.Printf("%p\n", slice5)
	fmt.Printf("%p", slice4)
	slice4 = append(slice4, 4, 4)
	slice5 = append(slice5, 5, 5)
	fmt.Println(slice4)
	fmt.Println(slice5)
}
