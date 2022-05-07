package main

import "fmt"

func If(ok int, a, b int) int
func LoopAdd(cnt, v0, step int) int
func main() {
	a, b := 2, 3
	value := If(0, a, b)
	fmt.Println(value)
}
