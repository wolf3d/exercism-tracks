package armstrong

import "math"
import "strconv"

func IsNumber(n int) bool {
	switch {
        case n >= 0 && n <= 9:
        	return true
        case n >= 10 && n < 100:
    		return false
        default:
    		return IsArmstrongNumber(n)
    }
	return false
}

func IsArmstrongNumber(n int) bool {
    digits := []rune(strconv.Itoa(n))
    length, sum :=len(digits), 0
    for i:=0; i < length; i++ {
        sum += int(math.Pow(float64(int(digits[i]-'0')),float64(length)))
    }
    return sum == n
}