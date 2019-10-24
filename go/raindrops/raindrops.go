package raindrops

import "strconv"

//Convert a number into a string that contains raindrop sounds
func Convert(input int) string {
	result := ""

	if input%3 == 0 {
		result += "Pling"
	}

	if input%5 == 0 {
		result += "Plang"
	}

	if input%7 == 0 {
		result += "Plong"
	}

	if result == "" {
		return strconv.Itoa(int(input))
	}

	return result
}
