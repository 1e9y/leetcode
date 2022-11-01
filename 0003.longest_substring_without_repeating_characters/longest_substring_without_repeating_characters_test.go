package longest_substring_without_repeating_characters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	input  string
	expect int
}{
	{
		"abcabcbb",
		3,
	},
	{
		"bbbbb",
		1,
	},
	{
		"pwwkew",
		3,
	},
	{
		"",
		0,
	},
}

func Test(t *testing.T) {
	for _, tt := range tests {
		assert.Equal(t, tt.expect, lengthOfLongestSubstring(tt.input))
	}
}
