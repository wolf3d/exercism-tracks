package parsinglogfiles

import "regexp"
import "strings"
func IsValidLine(text string) bool {
	re, err := regexp.Compile(`^(\[TRC\]|\[DBG\]|\[INF\]|\[WRN\]|\[ERR\]|\[FTL\])`)
    if err == nil {
    	return re.MatchString(text)
    }
    return false
}

func SplitLogLine(text string) []string {
	re, _ := regexp.Compile(`(\<[~*=-]{1,6}\>|<>)`)
    return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re, _ := regexp.Compile(`(?i)'.*password'|".*password.*"`)
    var count int = 0
    for _, v := range lines {
        f := re.FindAllString(v, -1)
        count += len(f)
    }
	return count
}

func RemoveEndOfLineText(text string) string {
	re, _ := regexp.Compile(`end-of-line[0-9]{1,}`)
    return re.ReplaceAllString(text, "")    
}

func TagWithUserName(lines []string) []string {
	re, _ := regexp.Compile(`User[ ]{1,}.*[0-9]`)    
    for i, v := range lines {
        if f := re.FindString(v); f != "" {
            userName := strings.Replace(f, "User", "", 1)
            userName = strings.Trim(userName," ")
            lines[i] = "[USR] " + userName + " " + v
        }        
    }
	return lines                      
}
