package main

import "fmt"

func main() {
	a, b, c, d, e, f, g := 0, 1, 2, 3, 4, 5, 6
	answer := doing(a, b, c, d, e, f, g)
	fmt.Println(answer)
}
func doing(a, b, c, d, e, f, g int) int {
	return a + b + c + d + e + f + g
}
