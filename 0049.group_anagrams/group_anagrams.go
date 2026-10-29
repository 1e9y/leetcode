package group_anagrams

import (
	"slices"
)

func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string, len(strs))

	for _, s := range strs {
		k := key(s)
		groups[k] = append(groups[k], s)
	}

	answer := make([][]string, 0, len(groups))
	for _, v := range groups {
		answer = append(answer, v)
	}

	return answer
}

func key(str string) string {
	s := []byte(str)
	slices.Sort(s)
	return string(s)
}
