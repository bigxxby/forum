package validation

func IsValidPost(title, content string) bool {
	if len(title) <= 0 || len(title) > 100 {
		return false
	}
	if len(content) <= 0 || len(content) >= 5000 {
		return false
	}
	return true

}
