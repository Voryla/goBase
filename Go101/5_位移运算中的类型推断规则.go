package main

func main() {
}

// 位移运算中的做类型不确定操作数类型的类型推断规则取决于右操作数是否是常量

const M = 2

var _ = 1.0 << M // 编译成功，1.0将被推断为int

var N = 2
var _ = 1.0 << N // 编译失败，1.0将被推断为float64
