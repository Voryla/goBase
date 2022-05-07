package main

import "fmt"

// conversion from int to string yields a string of one rune, not a string of digits (did you mean fmt.Sprint(x)?)
// UTF-8 编码中，十进制数字 65 对应的符号是 A。
func main() {
	i := 65

	fmt.Println(string(i))
}
