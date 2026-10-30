package valid_palindrome

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	l, r := 0, len(s)-1

	for l < r {
		a, b := rune(s[l]), rune(s[r])
		if !(unicode.IsLetter(a) || unicode.IsDigit(a)) {
			l++
			continue
		}
		if !(unicode.IsLetter(b) || unicode.IsDigit(b)) {
			r--
			continue
		}
		if a != b {
			return false
		}
		l++
		r--
	}

	return true
}
