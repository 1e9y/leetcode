package valid_palindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	input  string
	expect bool
}{
	{"A man, a plan, a canal: Panama", true},
	{"race a car", false},
	{" ", true},
}

func Test(t *testing.T) {
	for _, tt := range tests {
		assert.Equal(t, tt.expect, isPalindrome(tt.input))
	}
}
