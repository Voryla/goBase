package aboutQueue

func main() {
}

type MyQueue struct {
	data []int
}

func Constructor_1() MyQueue {
	myQueue := MyQueue{data: make([]int, 0)}
	return myQueue
}

func (this *MyQueue) Push(x int) {
	this.data = append(this.data, x)
}

func (this *MyQueue) Pop() int {
	defer func() { this.data = this.data[1:] }()
	return this.data[0]
}

func (this *MyQueue) Peek() int {
	return this.data[0]
}

func (this *MyQueue) Empty() bool {
	return len(this.data) == 0
}
