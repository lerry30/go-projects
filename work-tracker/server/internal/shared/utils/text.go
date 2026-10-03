package utils

import "strings"

func Capitalize(str string) string {
	str = strings.TrimSpace(str)
	if str == "" {
		return str
	}
	str = strings.ToLower(str)
	return strings.ToUpper(str[:1]) + str[1:]
}