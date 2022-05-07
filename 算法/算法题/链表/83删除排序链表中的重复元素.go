package aboutListNode

import "math"

func main() {
	answer := new(ListNode)
	answer.Val = 1
	answer.Next = &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, nil}}}}
	deleteDuplicates(answer)
}
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	answer := new(ListNode)
	answerHead := answer
	temp := math.MinInt
	var preNode *ListNode
	for head != nil {
		if head.Val > temp {
			answer.Val = head.Val
			temp = head.Val
			preNode = answer
			if head.Next != nil {
				answer.Next = &ListNode{0, nil}
				answer = answer.Next
			}
		}
		head = head.Next
	}
	if answer.Val == 0 && answer.Next == nil {
		preNode.Next = nil
	}
	return answerHead
}
