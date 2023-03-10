package climbing_stairs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	input  int
	expect int
}{
	{1, 1},
	{2, 2},
	{3, 3},
	{4, 5},
	{5, 8},
}

func Test(t *testing.T) {
	for _, tt := range tests {
		assert.Equal(t, tt.expect, climbStairs(tt.input))
	}
}
