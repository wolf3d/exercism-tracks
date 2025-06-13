// Package acronym provides functionality to create acronyms
package acronym

import "strings"

// Abbreviate creates acronim
func Abbreviate(s string) string {
    var abbr strings.Builder  
    re := strings.NewReplacer("_", "", "-", " ")
    cleaned := re.Replace(s)
	runes := []rune(cleaned)
    for i, val := range runes {
        if i == 0 {
            abbr.WriteString(strings.ToUpper(string(runes[0])))
        } else if val == ' ' && runes[i+1] != ' ' {
        	abbr.WriteString(strings.ToUpper(string(runes[i+1])))
        }        
    }
	return abbr.String()
}
