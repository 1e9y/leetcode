package top_k_frequent_elements

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	nums   []int
	k      int
	expect []int
}{
	{[]int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
	{[]int{1}, 1, []int{1}},
	{[]int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2}, 2, []int{1, 2}},
}

func Test(t *testing.T) {
	for _, tt := range tests {
		// The answer may be returned in any order.
		assert.ElementsMatch(t, tt.expect, topKFrequent(tt.nums, tt.k))
	}
}
