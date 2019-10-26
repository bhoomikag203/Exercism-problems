package accumulate

//Accumulate operation
func Accumulate(input []string, converter func(string) string) []string {

	output := make([]string, len(input))
	for i, s := range input {
		output[i] = converter(s)
	}
	return output
}
