package house_robber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	input  []int
	expect int
}{
	{[]int{1, 2, 3, 1}, 4},
	{[]int{2, 7, 9, 3, 1}, 12},
	{[]int{2, 1, 1, 2}, 4},
}

func Test(t *testing.T) {
	for _, tt := range tests {
		assert.Equal(t, tt.expect, rob(tt.input))
	}
}
