// Package leap contains method(s) to find out whether 
//year is leap year or not
package leap

// IsLeapYear reports if a year is leap year or not
func IsLeapYear(year int) bool {
	if year % 4 == 0 && year % 100 != 0 || year % 400 == 0 {
        return true
    }
	return false
}
