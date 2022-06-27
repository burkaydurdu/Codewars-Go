package weird_string_case

import (
	"strings"
	"unicode"
)

func ToWeirdCase(str string) string {
	var content string
	isNextUp := true

	for index, ch := range str {
		if isNextUp {
			content += strings.ToUpper(string(ch))
		} else {
			content += strings.ToLower(string(ch))
		}

		if string(ch) == " " && isNextUp {
			continue
		}

		if string(ch) == " " && !isNextUp {
			isNextUp = true
			continue
		}
		isNextUp = unicode.IsLower(rune(content[index]))
	}

	return content
}
