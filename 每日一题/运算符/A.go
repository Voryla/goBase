package main

/*
对于整数操作数，一元运算符 + , - 和 ^ 有如下定义：（省略了 ^ 的解释）
+x 　　　　 是 0 + x
-x 取其负值 是 0 - x
*/
func main() {
	v := new(int)
	*v = 2
	println(5 / +-*v)

}
