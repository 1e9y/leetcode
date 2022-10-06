package unique_paths

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testCases = []struct {
	input  []int
	expect int
}{
	{
		[]int{3, 2},
		3,
	},
	{
		[]int{3, 7},
		28,
	},
	{
		[]int{51, 9},
		1916797311,
	},
}

func TestUniquePaths(t *testing.T) {
	for _, c := range testCases {
		assert.Equal(t, c.expect, uniquePaths(c.input[0], c.input[1]))
	}
}
