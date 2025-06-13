package logs

import "unicode/utf8"
//import "fmt"
// Application identifies the application emitting the given log.
func Application(log string) string {
    const def string = "default"
    runes := []rune(log)

    for _, v := range runes {
        if v == '❗' || v == '🔍' || v == '☀' {
           switch v {
    			case '❗': return "recommendation"
            	case '🔍': return "search"
           		case '☀': return "weather"
           } 
        }
    }
    
	return def
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	runes := []rune(log)
    for i, v := range runes {
        if v == oldRune {
            runes[i] = newRune
        }
    }
return string(runes)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
