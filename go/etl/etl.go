package etl

import (
	"strings"
)

//Transform ..
func Transform(input map[int][]string) map[string]int {
	var c = []string{}
	output := make(map[string]int)
	for k := range input {
		c = input[k]
		for i := range c {
			c[i] = strings.ToLower(c[i])
			output[c[i]] = k
		}
	}
	return output
}
