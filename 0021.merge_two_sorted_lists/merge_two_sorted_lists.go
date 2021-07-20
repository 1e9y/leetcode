package merge_two_sorted_lists

import (
	"github.com/1e9y/leetcode/structures"
)

type ListNode = structures.ListNode

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	end := head
	var v int
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			v = l1.Val
			l1 = l1.Next
		} else {
			v = l2.Val
			l2 = l2.Next
		}

		end.Next = &ListNode{Val: v}
		end = end.Next
	}

	for l1 != nil {
		end.Next = &ListNode{Val: l1.Val}
		end = end.Next
		l1 = l1.Next
	}

	for l2 != nil {
		end.Next = &ListNode{Val: l2.Val}
		end = end.Next
		l2 = l2.Next
	}

	return head.Next
}
