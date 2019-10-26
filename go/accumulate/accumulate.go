package accumulate

//Accumulate operation
func Accumulate(input []string, converter func(string) string) (output []string) {
	for _, s := range input {
		output = append(output, converter(s))
	}
	return output
}
