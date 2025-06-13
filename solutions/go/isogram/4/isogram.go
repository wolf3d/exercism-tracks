package isogram

import (
    "unicode"
    )

func IsIsogram(word string) bool {
	runeSlice := []rune(word)
	mySet := make(map[rune]struct{})

	for i, v := range runeSlice {
		if v == ' ' || v == '-' {
			continue
		}
		if unicode.IsUpper(v) && unicode.IsLetter(v) {
			runeSlice[i] = unicode.ToLower(v)
		}
		if _, ok := mySet[runeSlice[i]]; ok {
			return false
		} else {
			mySet[runeSlice[i]] = struct{}{}
		}
	}
	return true
}
