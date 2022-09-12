package increasing_subsequences

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	nums   []int
	expect [][]int
}{
	{
		[]int{4, 6, 7, 7},
		[][]int{{4, 6}, {4, 6, 7}, {4, 6, 7, 7}, {4, 7}, {4, 7, 7}, {6, 7}, {6, 7, 7}, {7, 7}},
	},
	{
		[]int{4, 4, 3, 2, 1},
		[][]int{{4, 4}},
	},
}

func Test(t *testing.T) {
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			assert.ElementsMatch(t, tt.expect, findSubsequences(tt.nums))
		})
	}
}
