package etl

import (
	"strings"
)

//Transform ..
func Transform(input map[int][]string) map[string]int {
	output := make(map[string]int)
	for score, letters := range input {
		for _, letter := range letters {
			output[strings.ToLower(letter)] = score
		}
	}
	return output
}
