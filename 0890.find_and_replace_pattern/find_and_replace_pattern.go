package find_and_replace_pattern

func match(word, pattern string) bool {
	if len(word) != len(pattern) {
		return false
	}

	word2pattern := make(map[rune]bool)
	pattern2word := make(map[rune]rune)
	for i, w := range word {
		p := rune(pattern[i])
		if c, ok := pattern2word[p]; ok {
			if c != w {
				return false
			}
		} else {
			if word2pattern[w] {
				return false
			}
			word2pattern[w] = true
			pattern2word[p] = w
		}
	}
	return true
}

func findAndReplacePattern(words []string, pattern string) []string {
	result := make([]string, 0)

	for _, word := range words {
		if match(word, pattern) {
			result = append(result, word)
		}
	}

	return result
}
