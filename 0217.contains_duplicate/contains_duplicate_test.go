package contains_duplicate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	input  []int
	expect bool
}{
	{[]int{1, 2, 3, 1}, true},
	{[]int{1, 2, 3, 4}, false},
	{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
}

func Test(t *testing.T) {
	for _, tt := range tests {
		assert.Equal(t, tt.expect, containsDuplicate(tt.input))
	}
}
