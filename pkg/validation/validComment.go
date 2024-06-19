package validation

import "strings"

func IsValidComment(content string) bool {
	minLength := 1
	maxLength := 500

	content = strings.TrimSpace(content)
	if len(content) < minLength || len(content) > maxLength {
		return false
	}

	return true
}
