package main

import (
	"fmt"
)

type Math struct {
	x, y int
}

var m = map[string]Math{
	"foo": Math{2, 3},
}

// 解决办法1：修改数据结构，将Value修改为指针类型
//var m = map[string]*Math{
//	"foo": &Math{2, 3},
//}
func main() {
	// 错误原因：对于类似 X = Y的赋值操作，必须知道 X 的地址，才能够将 Y 的值赋给 X，但 go 中的 map 的 value 本身是不可寻址的。
	//m["foo"].x = 4
	// 解决办法1 ：使用临时变量
	temp := m["foo"]
	temp.x = 4
	m["foo"] = temp
	fmt.Println(m["foo"].x)
}
