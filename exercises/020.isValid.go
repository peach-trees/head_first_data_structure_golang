package exercises

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}

	checkMap := map[byte]byte{')': '(', ']': '[', '}': '{'}
	checkStack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			checkStack = append(checkStack, s[i])
		} else {
			if len(checkStack) == 0 {
				return false
			}
			if checkMap[s[i]] != checkStack[len(checkStack)-1] {
				return false
			}
			checkStack = checkStack[:len(checkStack)-1]
		} // else>>
	} // for>

	return len(checkStack) == 0
}
