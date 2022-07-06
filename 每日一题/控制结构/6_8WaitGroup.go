package main

import "sync"

// 下列代码会有什么问题？
// answer: 输出结果不唯一，代码存在风险，并非所有Go语句都能执行到
// 这是因为 go 执行太快了，导致 wg.Add(1) 还没有执行 main 函数就执行完毕了。wg.Add 的位置放错了。

const N = 10

var wg = &sync.WaitGroup{}

func main() {
	// 错误
	for i := 0; i < N; i++ {
		go func(i int) {
			wg.Add(1) // 可能导致main函数快速执行完毕
			println(i)
			defer wg.Done()
		}(i)
	}
	// 正确
	for i := 0; i < N; i++ {
		wg.Add(1)
		go func(i int) {
			println(i)
			defer wg.Done()
		}(i)
	}
	wg.Wait()
}
