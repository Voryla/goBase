package _defer

import "fmt"

// 参考答案及解析： F D M。
//被调用函数里的 defer 语句在返回之前就会被执行，所以输出顺序是 F D M。
func f() {
	defer fmt.Println("D")
	fmt.Println("F")
}

func main() {
	f()
	fmt.Println("M")
}
