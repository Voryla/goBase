package aboutStack

import (
	"fmt"
	"math"
)

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{math.MaxInt64},
	}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	top := this.minStack[len(this.minStack)-1]
	this.minStack = append(this.minStack, min(x, top))
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
func main() {

	obj := Constructor()
	obj.Push(-10)
	obj.Push(14)
	println(obj.GetMin())
	println(obj.GetMin())
	obj.Push(-20)
	println(obj.GetMin())
	println(obj.GetMin())
	println(obj.Top())
	println(obj.GetMin())
	obj.Pop()
	obj.Push(10)
	obj.Push(-7)
	println(obj.GetMin())
	obj.Push(-7)
	obj.Pop()
	println(obj.Top())
	fmt.Println(obj.GetMin())
	obj.Pop()
}
