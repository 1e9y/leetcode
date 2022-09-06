package find_and_replace_pattern

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	words   []string
	pattern string
	expect  []string
}{
	{
		[]string{"abc", "deq", "mee", "aqq", "dkd", "ccc"},
		"abb",
		[]string{"mee", "aqq"},
	},
	{

		[]string{"a", "b", "c"},
		"a",
		[]string{"a", "b", "c"},
	},
}

func Test(t *testing.T) {
	for _, tt := range tests {
		assert.Equal(t, tt.expect, findAndReplacePattern(tt.words, tt.pattern))
	}
}
