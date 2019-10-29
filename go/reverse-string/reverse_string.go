package reverse

//Reverse the given string
func Reverse(s string) string {
	// if s == "" {
	// 	return s
	// }
	// s1 := ""
	// for i := len(s) - 1; i >= 0; i-- {
	// 	s1 = s1 + string(s[i])
	// }
	// return s1
	res := make([]byte, len(s))
	prevPos, resPos := 0, len(s)
	for pos := range s {
		resPos -= pos - prevPos
		copy(res[resPos:], s[prevPos:pos])
		prevPos = pos
	}
	copy(res[0:], s[prevPos:])
	return string(res)
}
