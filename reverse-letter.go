package leetcodeplayground

// passed needs to optimize for memory
// 115 / 115 test cases passed.
// Status: Accepted
// Runtime: 0 ms
// Memory Usage: 2 MB
func ReverseOnlyLetters(s string) string {
	// 65 to 90 (a-z) and 97 to 122 (A-Z).
	resultString := make([]rune, len(s))
	sarr := []rune(s)
	// "j-Ih-gfE-dCba"
	for i, j := 0, len(sarr)-1; i < len(s); {
		if (sarr[i] >= 65 && sarr[i] <= 90) || (sarr[i] >= 97 && sarr[i] <= 122) {
			if (sarr[j] >= 65 && sarr[j] <= 90) || (sarr[j] >= 97 && sarr[j] <= 122) {
				resultString[i] = sarr[j]
				i++
			}
			j--
		} else {
			resultString[i] = sarr[i]
			i++
		}
	}
	return string(resultString)
}
