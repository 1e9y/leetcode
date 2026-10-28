package fibonacci_number

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	input  int
	expect int
}{
	{2, 1},
	{3, 2},
	{4, 3},
	{44, 300},
}

func Test(t *testing.T) {
	for _, tt := range tests {
		assert.Equal(t, tt.expect, fib(tt.input))
	}
}
