package checkout

import "regexp"

// ValidateUsername 对用户名的规则校验
func ValidateUsername(username string) bool {
	pattern := `^[a-zA-Z][a-zA-Z0-9_]{4,15}$`
	matched, _ := regexp.MatchString(pattern, username)
	return matched
}

// ValidatePassword 对密码进行规则校验
func ValidatePassword(password string) bool {
	pattern := `^[a-zA-Z]+[_.+*a-zA-Z0-9]{7,19}$`
	matched, _ := regexp.MatchString(pattern, password)
	return matched
}

// ValidateEmail 对邮箱进行规则校验
func ValidateEmail(email string) bool {
	pattern := `^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`
	matched, _ := regexp.MatchString(pattern, email)
	return matched
}
