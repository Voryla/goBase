package main

type sumifier interface {
	Add(a, b int32) int32
}
type sumer struct {
	id int32
}

func (math sumer) Add(a, b int32) int32 {
	return a + b
}

func main() {
	c := get()
	(*c).Add(1, 2)
}
func get() *sumifier {
	adder := sumer{id: 6754}
	m := sumifier(adder)
	return &m
}
