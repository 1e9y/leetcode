package add_two_numbers

// ListNode is a pre-defined singly-linked list type
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	curr := head
	mem := 0
	for l1 != nil || l2 != nil {
		val := mem
		if l1 != nil {
			val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val += l2.Val
			l2 = l2.Next
		}
		mem = val / 10
		val = val % 10
		next := &ListNode{
			Val:  val,
			Next: nil,
		}
		curr.Next = next
		curr = curr.Next
	}
	if mem != 0 {
		curr.Next = &ListNode{
			Val:  mem,
			Next: nil,
		}
	}
	return head.Next
}
