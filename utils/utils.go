package utils

import (
	"regexp"
	"strings"
)


func RemoveQuotes(str string) string {
	return strings.Trim(str, " \"")
}

func Sanitize(command string) string {
	var result []byte
	validChars := regexp.MustCompile("[a-zA-Z0-9=*+.\\-_/\\s]+")
	result = validChars.Find([]byte(command))
	return string(result)
}
