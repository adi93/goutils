package utils

import (
	"unicode"
)

// IsBlank - Returns true if a string is empty or is just whitespace characters
func IsBlank(s string) bool {
	if s == "" {
		return true
	}
	for _, c := range s {
		if unicode.IsSpace(c) {
			return false
		}
	}
	return true

}
