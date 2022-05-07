package main

func main() {
	s := Constructor()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	println(s.Pop())
	println(s.Pop())
	println(s.Pop())
}

type MyStack struct {
	Elem []int
}

/** Initialize your data structure here. */

func Constructor() MyStack {
	return MyStack{
		Elem: make([]int, 0),
	}
}

/** Push element x onto stack. */

func (stack *MyStack) Push(x int) {
	stack.Elem = append(stack.Elem, x)
}

/** Removes the element on top of the stack and returns that element. */

func (stack *MyStack) Pop() int {
	if stack.Empty() {
		return 0
	}
	x := stack.Elem[len(stack.Elem)-1]
	stack.Elem = stack.Elem[:(len(stack.Elem) - 1)]
	return x
}

/** Get the top element. */

func (stack *MyStack) Top() int {
	if stack.Empty() {
		return 0
	}
	return stack.Elem[len(stack.Elem)-1]
}

/** Returns whether the stack is empty. */

func (stack *MyStack) Empty() bool {
	return len(stack.Elem) == 0
}
