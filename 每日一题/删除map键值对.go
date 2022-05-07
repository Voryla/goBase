package main

import "fmt"

// 删除map不存在的值，不会报错，相当于没有任何作用；获取不存在的值时，返回值类型对应的零值，所以返回0
func main() {
	s := make(map[string]int)
	delete(s, "h")
	fmt.Println(s["h"])
}
