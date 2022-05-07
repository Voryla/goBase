package main

import "fmt"

func main() {
	// 1. switch 流程控制块中的case表达式不能重复，但是布尔值可以重复
	//switch 123 {
	//case 123:
	//case 123: // error: duplicate case
	//}
	// 但是下面的代码在编译时是没问题的。
	//switch false {
	//case false:
	//case false:
	//}

	// 2. 当switch表达式省略时，其默认值为bool的确定值 true，下面代码将永远输出 true
	switch { // <=> switch true {
	case true:
		fmt.Println("true")
	case false:
		fmt.Println("false")
	}

	// 3. switch 表达式总是被估值为类型确定值
	// 例如：下列代码块中的switch表达式123 被视为一个int值，而不是一个类型不确定的整数
	switch 123 {
	//case int64(123):
	//case uint32(123):
	}

	// 4. 若switch 左大括放置在switch下方，由于Go代码的断行规则，switch False()后将被插入; 那么下列代码运行将会输出true
	//switch False()
	//{
	//case true:  fmt.Println("true")
	//case false: fmt.Println("false")
	//}
}
func False() bool {
	return false
}
