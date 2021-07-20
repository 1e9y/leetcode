package merge_two_sorted_lists

import (
	"fmt"
	"testing"

	"github.com/1e9y/leetcode/structures"
	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	input  [][]int
	expect []int
}{
	{
		[][]int{{1, 2, 4}, {1, 3, 4}},
		[]int{1, 1, 2, 3, 4, 4},
	},
	{
		[][]int{{}, {}},
		[]int{},
	},
	{
		[][]int{{}, {0}},
		[]int{0},
	},
	{
		[][]int{{1}, {}},
		[]int{1},
	},
	{
		[][]int{{1, 1, 1, 1}, {2, 2, 2, 2}},
		[]int{1, 1, 1, 1, 2, 2, 2, 2},
	},
}

func TestMergeTwoLists(t *testing.T) {
	for _, c := range testCases {
		t.Run(fmt.Sprintf("%v+%v", c.input[0], c.input[1]), func(t *testing.T) {
			a, b := structures.FromInts(c.input[0]), structures.FromInts(c.input[1])
			e := structures.FromInts(c.expect)
			m := mergeTwoLists(a, b)
			assert.True(t, e.Equal(m), "got: %v\nwant: %v", m, e)
		})
	}
}
