package valid_anagram

import "slices"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	a, b := []byte(s), []byte(t)
	slices.Sort(a)
	slices.Sort(b)

	return slices.Equal(a, b)
}
