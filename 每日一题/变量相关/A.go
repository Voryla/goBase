package main

import "fmt"

// 因为循环体内的val为同一个指针类型(处于同一个地址,只是循环时不断改变值). 因此, &val是固定的. 循环结束时val == 3, 因此最后所有的*v == 3
func main() {
	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	for key, val := range slice {
		m[key] = &val
	}

	for k, v := range m {
		fmt.Println(k, "->", *v)
	}
}
