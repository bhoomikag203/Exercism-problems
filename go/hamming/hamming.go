package hamming

import (
	"errors"
)

//Distance function calculates the hamming distance
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("the length of the strings must be equal")
	}

	a1 := []rune(a)
	b1 := []rune(b)
	count := 0
	for i := 0; i < len(a1); i++ {
		if a1[i] != b1[i] {
			count++
		}
	}
	return count, nil
}
