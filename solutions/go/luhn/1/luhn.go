package luhn
import ( 
    //"fmt"
    "unicode"
    )
func Valid(id string) bool {
	var double bool = false    
    var digits []rune = []rune(id)
    var length int = len(digits)
    var sum int = 0
    if length == 1 && digits[0] == '0' {
        return false
    }
    for i := length-1; i >= 0; i-- {
        var v rune = digits[i]
        if unicode.IsDigit(v) {
            var digit int = int(v - '0')
            if double {
                digit *= 2
                if digit > 9 {
                    digit -= 9
                }
                double = false
            } else {
            	double = true
            }        
            sum += digit
        } else if v == ' ' && length == 2 {
        	return false
        } else if v == ' ' {
			continue
		} else {
        	return false
        }
    }
    return sum % 10 == 0
}