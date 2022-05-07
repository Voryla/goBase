package aboutListNode

import "fmt"

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev

}

//func reverseList(head *ListNode) *ListNode {
//	var prev *ListNode
//	curr := head
//	for curr != nil {
//		// 先取出next
//		next := curr.Next
//		// 将当前链表元素的next替换为倒叙链表
//		curr.Next = prev
//		// 将当前元素放置在倒叙链表的首部
//		prev = curr
//		// 遍历下一个元素
//		curr = next
//	}
//	return prev
//}

func reverseList1(head *ListNode) *ListNode {
	//
	if head == nil || head.Next == nil {
		return head
	}
	//
	newHead := reverseList1(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

func main() {
	s := &ListNode{
		Val:  0,
		Next: new(ListNode),
	}

	s.Next.Val = 1
	s1 := new(ListNode)
	s1.Val = 2
	s.Next.Next = s1

	//for s != nil {
	//	fmt.Println(s.Val)
	//	s = s.Next
	//}

	s = reverseList(s)
	for s != nil {
		fmt.Println(s.Val)
		s = s.Next
	}
}
