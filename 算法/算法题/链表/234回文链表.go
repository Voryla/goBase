package aboutListNode

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := new(ListNode)
	head.Val = 1
	head.Next = &ListNode{2, &ListNode{2, &ListNode{1, nil}}}
	println(isPalindrome(head))
}
func isPalindrome(head *ListNode) bool {
	values := make([]int, 0)
	for head != nil {
		values = append(values, head.Val)
		head = head.Next
	}
	count := len(values)
	for i, v := range values[:count/2] {
		if v != values[count-1-i] {
			return false
		}
	}
	return true
}
