package valid_anagram

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	s      string
	t      string
	expect bool
}{
	{"anagram", "nagaram", true},
	{"rat", "car", false},
}

func Test(t *testing.T) {
	for _, tt := range tests {
		assert.Equal(t, tt.expect, isAnagram(tt.s, tt.t))
	}
}
