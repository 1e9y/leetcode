package group_anagrams

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	input  []string
	expect [][]string
}{
	{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}},
	{[]string{""}, [][]string{{""}}},
	{[]string{"a"}, [][]string{{"a"}}},
}

// sortGroups sorts strings inside each group, since the answer may come in any order.
func sortGroups(groups [][]string) [][]string {
	sorted := make([][]string, len(groups))
	for i, g := range groups {
		sorted[i] = slices.Clone(g)
		slices.Sort(sorted[i])
	}
	return sorted
}

func Test(t *testing.T) {
	for _, tt := range tests {
		assert.ElementsMatch(t, sortGroups(tt.expect), sortGroups(groupAnagrams(tt.input)))
	}
}
