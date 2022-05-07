package main

import (
	"fmt"
	"time"
)

// 一个声明了类型但没有显式的字段名的字段就是 嵌入字段 。嵌入字段必须指定为一个类型名 T 或者为一个到非接口类型的指针名 *T ，并且 T 不是一个指针类型。这个非限定的类型名就被当作字段名。
// 注意：一种类型只能出现一次，因为受制于：在结构体中，非 空白 字段名必须是 唯一的 。
type noName struct {
	int
	float32
}

func main() {
	c := noName{
		int:     3,
		float32: 4,
	}
	fmt.Println(c)
	switch { // <=> switch true {
	case c.int > 4:
		fmt.Println("true")
	case c.int > 5:
		fmt.Println("false")
	}
	switch alwaysFalse(); {
	case true:
		fmt.Println("true")
	case false:
		fmt.Println("false")
	}

	newStruct := new(noName)
	fmt.Printf("%p\n", newStruct)
	fmt.Printf("%p\n", &newStruct)
	fmt.Printf("%v\n", newStruct)

	newStruct1 := new(noName)
	fmt.Printf("%p\n", newStruct1)
	fmt.Printf("%p\n", &newStruct1)
	fmt.Printf("%v\n", newStruct1)

	loc := time.FixedZone("UTC+8", +8*60*60)
	time1 := time.Date(2022, 3, 4, 5, 3, 12, 3, loc)
	z := time1.Format("2006-1-2 15:04:5")
	fmt.Println(z)
}

func alwaysFalse() bool {
	return true
}
