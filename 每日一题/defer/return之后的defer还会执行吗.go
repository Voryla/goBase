package _defer

import "fmt"

// defer 若想执行，必须注册在return之前
var a bool = true

func main() {
	defer func() {
		fmt.Println("1")
	}()
	if a == true {
		fmt.Println("2")
		return
	}
	defer func() {
		fmt.Println("3")
	}()
}
