package isbn

import "regexp"

// BenchmarkIsValidISBN-8              9146            131067 ns/op          110912 B/op       1037 allocs/op
// performance is ridiculously bad :D
func IsValidISBN(isbn string) bool {
	var length int = len(isbn)
	re := regexp.MustCompile("[0-9]{1}-[0-9]{3}-[0-9]{5}-[0-9xX]|[0-9]{9}[0-9xX]{1}")
	switch length {
	case 10:
		return ValidateIsbn(re, isbn, length)
	case 13:
		return ValidateIsbn(re, isbn, length)
	default:
		return false
	}
	return false
}

func ValidateIsbn(re *regexp.Regexp, isbn string, length int) bool {
	isbnSum := 0
	if re.MatchString(isbn) {
		if isbn[length-1] == 'X' {
			isbnSum = 10
		} else {
			isbnSum = int(isbn[length-1] - '0')
		}

		isbnSum += SumIsbn(isbn, length)

		return isbnSum%11 == 0
	}
	return false
}

func SumIsbn(isbn string, length int) int {
	switch length {
	case 10:
		return int(isbn[0]-'0')*10 + int(isbn[1]-'0')*9 + int(isbn[2]-'0')*8 + int(isbn[3]-'0')*7 + int(isbn[4]-'0')*6 + int(isbn[5]-'0')*5 + int(isbn[6]-'0')*4 + int(isbn[7]-'0')*3 + int(isbn[8]-'0')*2
	case 13:
		return int(isbn[0]-'0')*10 + int(isbn[2]-'0')*9 + int(isbn[3]-'0')*8 + int(isbn[4]-'0')*7 + int(isbn[6]-'0')*6 + int(isbn[7]-'0')*5 + int(isbn[8]-'0')*4 + int(isbn[9]-'0')*3 + int(isbn[10]-'0')*2
	}
	return 0
}
