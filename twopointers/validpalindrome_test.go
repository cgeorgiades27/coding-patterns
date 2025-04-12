package twopointers

import "testing"

func validPalindromes(str string) bool {
	left, right := 0, len(str)-1
	for left < right {
		for !isAlnum(str[left]) && left < right {
			left++
		}
		for !isAlnum(str[right]) && right > left {
			right--
		}
		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func isAlnum(c byte) bool {
	return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c >= '0' && c <= '9'
}

func TestValidPalindrome(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"!, (?)", true},
		{"12.02.2021", true},
		{"21.02.2021", false},
		{"hello, world!", false},
		{"a dog! a panic in a pagoda.", true},
	}
	for _, test := range tests {
		if test.want != validPalindromes(test.input) {
			t.Errorf("test failed, input:%v, got:%v", test.input, test.want)
		}
	}
}
