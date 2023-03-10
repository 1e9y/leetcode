package longest_increasing_subsequence

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	nums   []int
	expect int
}{
	{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
	{[]int{0, 1, 0, 3, 2, 3}, 4},
	{[]int{7, 7, 7, 7, 7, 7, 7}, 1},
	{[]int{1, 3, 6, 7, 9, 4, 10, 5, 6}, 6},
}

func Test(t *testing.T) {
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.nums), func(t *testing.T) {
			assert.Equal(t, tt.expect, lengthOfLIS(tt.nums))
		})
	}
}
