package sorting_the_sentence

import (
	"strings"
)

func sortSentence(s string) string {
	words := strings.Split(s, " ")
	output := make([]string, len(words))
	for _, w := range words {
		i := w[len(w)-1]
		output[i-'1'] = w[:len(w)-1]
	}
	return strings.Join(output, " ")
}
