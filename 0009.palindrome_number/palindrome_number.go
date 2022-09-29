package palindrome_number

// isPlaindrome returns whether a number is palindrome.
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	if x%10 == 0 {
		return false
	}
	var r int
	for x > r {
		r = r*10 + x%10
		x /= 10
	}
	return x == r || x == r/10
}
