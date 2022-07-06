package aboutListNode

// 空间换时间
func middleNode(head *ListNode) *ListNode {
	temp := make([]*ListNode, 0)
	for head != nil {
		temp = append(temp, head)
		head = head.Next
	}
	return temp[len(temp)/2]
}

// 快慢指针
func qsPointer(head *ListNode) *ListNode {
	p, q := head, head
	for q != nil && q.Next != nil {
		// 快指针每次走两步
		q = q.Next.Next
		// 慢指针每次走一步，当快指针行至末尾时，慢指针直行一半
		p = p.Next
	}
	return p
}
