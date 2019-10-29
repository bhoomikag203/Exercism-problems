package isogram

import (
	"strings"
)

//IsIsogram checks if the input is isogram
func IsIsogram(s string) bool {
	str := strings.ToLower(s)
	seen := map[rune]bool{}

	for _, r := range str {
		if r == '-' || r == ' ' {
			continue
		}
		if seen[r] {
			return false
		}
		seen[r] = true
	}
	return true
}
