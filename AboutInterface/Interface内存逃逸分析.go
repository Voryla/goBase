package main

type notifier interface {
	Add(a, b int32) int32
}
type Sumer struct {
	id int32
}

func (math Sumer) Add(a, b int32) int32 {
	return a + b
}

func main() {
	adder := Sumer{id: 123}
	m := notifier(adder)
	m.Add(1, 12)
}
