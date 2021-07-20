package structures

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func FromInts(nums []int) *ListNode {
	head := new(ListNode)
	end := head
	for _, n := range nums {
		end.Next = &ListNode{Val: n}
		end = end.Next
	}
	return head.Next
}

func (l *ListNode) Equal(m *ListNode) bool {
	for l != nil && m != nil {
		if l.Val != m.Val {
			return false
		}
		l, m = l.Next, m.Next
	}
	return l == nil && m == nil
}

func (l *ListNode) String() (s string) {
	for {
		s += fmt.Sprint(l.Val)
		if l.Next != nil {
			s += " -> "
			l = l.Next
		} else {
			return s
		}
	}
}
