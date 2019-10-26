package isogram

import (
	"strings"
)

//IsIsogram checks if the input is isogram
func IsIsogram(s string) bool {
	str := []rune(strings.ToLower(s))
	letterCountMap := map[rune]int{}

	for _, key := range str {
		if key == '-' || key == ' ' {
			continue
		}

		letterCountMap[key]++
		if letterCountMap[key] > 1 {
			return false
		}
	}
	return true
}
