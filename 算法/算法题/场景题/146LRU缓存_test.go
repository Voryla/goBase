package 场景题

import "testing"

func TestName(t *testing.T) {
	l := Constructor(1)
	l.Put(2, 1)
	l.Get(2)
	l.Put(3, 2)
	l.Get(2)
	l.Get(3)
}
