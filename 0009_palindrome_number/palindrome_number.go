package _0009_palindrome_number

// this method checks if the given number is palindrome or not
func isPalindrome(x int) bool {
	i := x
	reverse := 0
	for i > 0 {
		r := i % 10
		reverse = r + reverse*10
		i = i / 10
	}
	if reverse != x {
		return false
	}
	return true
}
