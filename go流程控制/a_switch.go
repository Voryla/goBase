package main

import "fmt"

func main() {
	var k int
	fmt.Scan(&k)
	switch k {
	case 3:
		fmt.Println(3)
		break // 显示添加break
	case 2:
		fmt.Println(2)
		fallthrough // 会执行后续的case分支
	case 1:
		fmt.Println(1) // go默认会添加break
	default:
		fmt.Println("default")
	}
}
