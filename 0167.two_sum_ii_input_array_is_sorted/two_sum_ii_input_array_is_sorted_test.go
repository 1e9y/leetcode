package two_sum_ii_input_array_is_sorted

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	numbers []int
	target  int
	expect  []int
}{
	{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
	{[]int{2, 3, 4}, 6, []int{1, 3}},
	{[]int{-1, 0}, -1, []int{1, 2}},
}

func Test(t *testing.T) {
	for _, tt := range tests {
		assert.Equal(t, tt.expect, twoSum(tt.numbers, tt.target))
	}
}
