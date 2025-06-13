package isogram

import "strings"

func IsIsogram(word string) bool {
    charMap := make(map[rune]int)
    var phrase string = strings.ReplaceAll(word,"-","")
    phrase = strings.ReplaceAll(phrase," ","")
	for _,v := range strings.ToLower(phrase) {
        _, exists := charMap[v]
        if exists {
            return false
        } else { 
        	charMap[v] += 1
        }
    }
	return true
}
