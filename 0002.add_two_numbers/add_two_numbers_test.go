package add_two_numbers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	input  [][]int
	expect []int
}{
	{
		[][]int{
			{2, 4, 3},
			{7, 0, 8}},
		[]int{7, 0, 8},
	},
	{
		[][]int{
			{9, 9, 9, 9, 9, 9, 9},
			{9, 9, 9, 9},
		},
		[]int{8, 9, 9, 9, 0, 0, 0, 1},
	},
}

func numsToList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	node := &ListNode{
		Val: nums[0],
	}
	for i := 1; i < len(nums); i++ {
		next := &ListNode{Val: nums[i]}
		node.Next = next
		node = next
	}

	return node
}

func Test(t *testing.T) {
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.input), func(t *testing.T) {
			assert.Equal(t, tt.expect, addTwoNumbers(
				numsToList(tt.input),
				numsToList(tt.expect),
			))
		})
	}
}
