package sum_of_all_subset_xor_totals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	nums   []int
	expect int
}{
	{
		[]int{1, 3},
		6,
	},
	{
		[]int{5, 1, 6},
		28,
	},
	{
		[]int{3, 4, 5, 6, 7, 8},
		480,
	},
}

func TestSubsetXORSum(t *testing.T) {
	for _, c := range testCases {
		assert.Equal(t, c.expect, subsetXORSum(c.nums))
	}
}
