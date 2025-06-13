package luhn

func Valid(id string) bool {
	var start int = len(id) - 1
	var counter, sum int = 0, 0
	for i := start; i > -1; i-- {
		character := id[i]
		digit := int(character - '0')
		switch {
		case character == ' ':
			continue
		case character >= '0' && character <= '9':
			if counter%2 == 0 {
				sum += digit
			} else {
				switch digit {
				case 1, 2, 3, 4:
					sum += digit << 1
				case 5:
					sum += 1
				case 6:
					sum += 3
				case 7:
					sum += 5
				case 8:
					sum += 7
				case 9:
					sum += digit
				}
			}
			counter += 1
		default:
			return false
		}
	}
	return counter > 1 && sum%10 == 0
}
