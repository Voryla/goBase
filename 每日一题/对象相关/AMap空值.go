package main

import "fmt"

type person struct {
	name string
}

// m 是一个 map，值是 nil。从 nil map 中取值不会报错，而是返回相应的零值，这里值是 int 类型，因此返回 0。
func main() {
	var m map[person]int
	p := person{"mike"}
	fmt.Println(m[p])
}
