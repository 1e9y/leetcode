package sorting_the_sentence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	s      string
	expect string
}{
	{
		"is2 sentence4 This1 a3",
		"This is a sentence",
	},
	{
		"Myself2 Me1 I4 and3",
		"Me Myself and I",
	},
}

func TestSortSentence(t *testing.T) {
	for _, c := range testCases {
		assert.Equal(t, c.expect, sortSentence(c.s))
	}
}
