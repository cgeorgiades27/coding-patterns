package validpalindrome

import (
	"testing"
)

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
