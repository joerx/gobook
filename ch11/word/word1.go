package word

import (
	"unicode"
)

func IsPalindrome(s string) bool {
	letters := make([]rune, 0, len(s))
	// letters := make([]rune, 0, 0)

	for _, l := range s {
		if unicode.IsLetter(l) {
			letters = append(letters, unicode.ToLower(l))
		}
	}

	n := len(letters)/2

	for i := 0; i < n; i++ {
		if letters[i] != letters[len(letters)-i-1] {
			return false
		}
	}
	return true
}
