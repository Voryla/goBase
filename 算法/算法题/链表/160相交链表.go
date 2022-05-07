package aboutListNode

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func main() {
	//a := &ListNode{Val: 8, Next: nil}
	b := &ListNode{Val: 4, Next: nil}
	b.Next = &ListNode{Val: 1}
	b.Next.Next = &ListNode{Val: 8}
	b.Next.Next.Next = &ListNode{Val: 4}
	b.Next.Next.Next.Next = &ListNode{Val: 5, Next: b.Next.Next}
	//println(getIntersectionNode(a, b).Val)
	println(hasCycle(b))
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	vis := map[*ListNode]bool{}
	for tmp := headA; tmp != nil; tmp = tmp.Next {
		vis[tmp] = true
	}
	for tmp := headB; tmp != nil; tmp = tmp.Next {
		if vis[tmp] {
			return tmp
		}
	}
	return nil
}
func hasCycle(head *ListNode) bool {
	tempMap := map[*ListNode]bool{}
	for temp := head; temp != nil; temp = temp.Next {
		if tempMap[temp.Next] == true {
			return true
		} else {
			tempMap[temp] = true
		}
	}
	return false
}
