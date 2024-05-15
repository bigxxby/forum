package validation

func IsValidComment(content string) bool {
	minLength := 1
	maxLength := 500

	if len(content) < minLength || len(content) > maxLength {
		return false
	}

	return true
}
