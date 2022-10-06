package two_sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	numbers []int
	target  int
	expect  []int
}{
	{
		[]int{2, 7, 11, 15},
		9,
		[]int{0, 1},
	},
	{
		[]int{3, 2, 4},
		6,
		[]int{1, 2},
	},
	{
		[]int{0, 1},
		1,
		[]int{0, 1},
	},
}

func TestTwoSums(t *testing.T) {
	for _, c := range testCases {
		assert.Equal(t, c.expect, twoSum(c.numbers, c.target))
	}
}
