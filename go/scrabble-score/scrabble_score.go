package scrabble

//Score function compute the scrabble score for a given word
func Score(input string) int {
	score := 0
	for _, v := range input {
		switch v {
		case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T', 'a', 'e', 'i', 'o', 'u', 'l', 'n', 'r', 's', 't':
			score++
		case 'D', 'G', 'd', 'g':
			score += 2
		case 'B', 'C', 'M', 'P', 'b', 'c', 'm', 'p':
			score += 3
		case 'F', 'H', 'V', 'W', 'Y', 'f', 'h', 'v', 'w', 'y':
			score += 4
		case 'K', 'k':
			score += 5
		case 'J', 'X', 'j', 'x':
			score += 8
		case 'Q', 'Z', 'q', 'z':
			score += 10
		}
	}
	return score
}
