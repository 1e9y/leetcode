package increasing_subsequences

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	input  int
	expect int
}{
    // TODO: Add test cases
}

func Test(t *testing.T) {
	for _, tt := range tests {
		assert.Equal(t, tt.expect, sol(tt.numbers, tt.target))
	}
}
