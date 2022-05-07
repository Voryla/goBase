package aboutListNode

func main() {
	b := &ListNode{Val: 4, Next: nil}
	b.Next = &ListNode{Val: 1}
	b.Next.Next = &ListNode{Val: 8}
	b.Next.Next.Next = &ListNode{Val: 4}
	b.Next.Next.Next.Next = &ListNode{Val: 5, Next: nil}
	//println(getIntersectionNode(a, b).Val)
	t := getKthFromEnd(b, 2)
	println(t.Val)
	println(t.Next.Val)
}
func getKthFromEnd(head *ListNode, k int) *ListNode {
	temp := 0
	list := head
	for list != nil {
		temp++
		list = list.Next
	}
	for i := 0; i < temp-k; i++ {
		head = head.Next
	}
	return head
}
