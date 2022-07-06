package aboutListNode

// problem： 给你一个链表的头节点 head 和一个整数 val ，请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 。
func removeElements(head *ListNode, val int) *ListNode {
	header := &ListNode{0, head}
	curr := header
	for curr.Next != nil {
		if curr.Next.Val == val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}
	return header.Next
}
