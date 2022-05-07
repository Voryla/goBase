package aboutListNode

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	top := &ListNode{}
	temp := top
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			temp.Val = list1.Val
			list1 = list1.Next
		} else {
			temp.Val = list2.Val
			list2 = list2.Next
		}
		temp.Next = &ListNode{}
		temp = temp.Next
	}
	if list1 != nil {
		temp.Val = list1.Val
		temp.Next = list1.Next
	} else if list2 != nil {
		temp.Val = list2.Val
		temp.Next = list2.Next
	} else {
		return nil
	}
	return top
}
