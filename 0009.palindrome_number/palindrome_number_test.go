package palindrome_number

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	input  int
	expect bool
}{
	{
		121,
		true,
	},
	{
		-121,
		false,
	},
	{
		10,
		false,
	},
	{
		123,
		false,
	},
}

func Test(t *testing.T) {
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.input), func(t *testing.T) {
			assert.Equal(t, tt.expect, isPalindrome(tt.input))
		})
	}
}
