package not_very_secure

import (
	"regexp"
)

func Alphanumeric(str string) bool {
	if len(str) == 0 {
		return false
	}
	re, _ := regexp.Compile("[^a-zA-Z\\d]")
	matchText := re.FindString(str)
	return len(matchText) == 0
}
