package isogram

import (
    "strings"
    "unicode"
    )


func IsIsogram(word string) bool {
	word = strings.ReplaceAll(word, " ", "")
	word = strings.ReplaceAll(word, "-", "")
	runeSlice := []rune(word)

	for i, v := range runeSlice {
		if unicode.IsUpper(v) && unicode.IsLetter(v) {
			runeSlice[i] = unicode.ToLower(v)
		}
	}
	lowerCasedWord := string(runeSlice)
	stringSlice := strings.Split(lowerCasedWord, "")

	for _, v := range stringSlice {
		// if v == " " || v == "-" {
		// 	continue
		// }
		if strings.Count(lowerCasedWord, v) > 1 {
			return false
		}
	}
	return true
}
