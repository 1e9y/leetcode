package structures

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromInts(t *testing.T) {
	l := &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}}
	n := FromInts([]int{0, 1, 2})
	assert.True(t, l.Equal(n))
}

func TestListNode_Equal(t *testing.T) {
	m := &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}}
	assert.True(t, m.Equal(m))

	n := &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}}
	assert.True(t, m.Equal(n))

	k := &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 3}}}
	assert.False(t, m.Equal(k))
}

func TestListNode_String(t *testing.T) {
	tests := []struct {
		ints   []int
		expect string
	}{
		{[]int{}, fmt.Sprint(nil)},
		{[]int{0}, "0"},
		{[]int{0, 1, 2}, "0 -> 1 -> 2"},
		{[]int{0, 1, 2, 3}, "0 -> 1 -> 2 -> 3"},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expect, fmt.Sprint(FromInts(tt.ints)))
	}
}
