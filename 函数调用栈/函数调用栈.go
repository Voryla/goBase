package main

import "unicode/utf8"

func mul(a int, b int) int {
	return a * b
}
func main() {
	num1, num2 := 1, 3
	c := mul(num1, num2)
	_ = c
	s, _ := utf8.DecodeRuneInString("")
	println(s)
}
