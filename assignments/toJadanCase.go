package assignments

import (
	"strings"
)

func ToJodanCase(str string) string {
	var result string
	words := strings.Split(str, " ")
	for index, word := range words {
		wordUpper := strings.ToUpper(word[0:1]) + word[1:]
		words[index] = wordUpper
	}
	result = strings.Join(words, " ")
	return result
}

// func ToJodanCaseBestPractices(str string) string {
// 	return strings.Title(str)
// }
