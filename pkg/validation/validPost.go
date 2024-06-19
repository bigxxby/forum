package validation

import "strings"

func IsValidPost(title, content string) bool {

	title = strings.TrimSpace(title)
	content = strings.TrimSpace(content)

	if len(title) <= 0 || len(title) > 100 {
		return false
	}
	if len(content) <= 0 || len(content) >= 5000 {
		return false
	}
	return true
}
