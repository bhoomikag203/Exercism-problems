package isogram

import (
	"unicode"
)

//IsIsogram checks if the input is isogram
func IsIsogram(s string) bool {
	seen := map[rune]bool{}

	for _, r := range s {
		if r == '-' || r == ' ' {
			continue
		}
		r = unicode.ToLower(r)
		if seen[r] {
			return false
		}
		seen[r] = true
	}
	return true
}
