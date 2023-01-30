package non_decreasing_subsequences

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	input  []int
	expect [][]int
}{
	{
		input:  []int{4, 6, 7, 7},
		expect: [][]int{{4, 6}, {4, 7}, {4, 6, 7}, {4, 6, 7, 7}, {6, 7}, {6, 7, 7}, {7, 7}, {4, 7, 7}},
	},
	{
		input:  []int{4, 4, 3, 2, 1},
		expect: [][]int{{4, 4}},
	},
}

func Test(t *testing.T) {
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.ElementsMatch(t, tt.expect, findSubsequences(tt.input))
		})
	}
}
