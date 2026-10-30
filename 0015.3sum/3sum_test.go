package three_sum

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	input  []int
	expect [][]int
}{
	{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
	{[]int{0, 1, 1}, [][]int{}},
	{[]int{0, 0, 0}, [][]int{{0, 0, 0}}},
	{[]int{-2, 0, 0, 2, 2}, [][]int{{-2, 0, 2}}},
}

// sortTriplets sorts numbers inside each triplet, since the answer may come in any order.
func sortTriplets(triplets [][]int) [][]int {
	sorted := make([][]int, len(triplets))
	for i, t := range triplets {
		sorted[i] = slices.Clone(t)
		slices.Sort(sorted[i])
	}
	return sorted
}

func Test(t *testing.T) {
	for _, tt := range tests {
		assert.ElementsMatch(t, sortTriplets(tt.expect), sortTriplets(threeSum(tt.input)))
	}
}
