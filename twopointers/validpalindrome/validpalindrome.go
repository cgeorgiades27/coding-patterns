package validpalindrome

import "strings"

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
