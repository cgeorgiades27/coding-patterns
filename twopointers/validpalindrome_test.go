package twopointers

import (
	"strings"
	"testing"
)

func validPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		for !isAlnum(s[left]) && left < right {
			left++
		}
		for !isAlnum(s[right]) && right > left {
			right--
		}
		strLeft := string(s[left])
		strRight := string(s[right])
		if strings.ToLower(strLeft) != strings.ToLower(strRight) {
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
		{"A man, a plan, a canal: Panama", true},
	}
	for _, test := range tests {
		if test.want != validPalindrome(test.input) {
			t.Errorf("test failed, input:%v, got:%v", test.input, test.want)
		}
	}
}
