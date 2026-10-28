package n_th_tribonacci_number

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	input  int
	expect int
}{
	{4, 4},
	{25, 1389537},
}

func Test(t *testing.T) {
	for _, tt := range tests {
		assert.Equal(t, tt.expect, tribonacci(tt.input))
	}
}
