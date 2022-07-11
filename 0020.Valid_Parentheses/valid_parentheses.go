package _020_Valid_Parentheses

// this method will return if the give parenthesis string is valid or not
func isValid(s string) bool {
	//stack as an array representation
	var stack []uint8
	for i := 0; i < len(s); i++ {
		n := len(stack) - 1
		if s[i] == ')' && n >= 0 {
			if stack[n] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		} else if s[i] == '}' && n >= 0 {
			if stack[n] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		} else if s[i] == ']' && n >= 0 {
			if stack[n] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
