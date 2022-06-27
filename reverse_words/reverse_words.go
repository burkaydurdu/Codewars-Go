package reverse_words

import "strings"

func reverse(str string) string {
	rns := []rune(str)

	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		rns[i], rns[j] = rns[j], rns[i]
	}

	return string(rns)
}

func ReverseWords(str string) string {
	stringArray := strings.Split(str, " ")

	for index, item := range stringArray {
		stringArray[index] = reverse(item)
	}

	return strings.Join(stringArray, " ")
}
