package utils

import (
	"strings"
)

func ExtractConfigValue(configLine string) string {
	idx := strings.Index(configLine, "=")
	result := configLine
	if idx > 0 {
		result = configLine[idx+1:]
	}
	return strings.Trim(result, " \"")
}

func RemoveQuotes(str string) string {
	return strings.Trim(str, "\"")
}