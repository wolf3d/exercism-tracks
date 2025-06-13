package isogram

func IsIsogram(word string) bool {
	phraseLength := len(word)
	mySet := make(map[int]struct{})

	for i := 0; i < phraseLength; i++ {
		character := int(word[i])

		if character >= 65 && character <= 90 {
			if OperateSet(mySet, 90-character) {
				return false
			}
		}
		if character >= 97 && character <= 122 {
			if OperateSet(mySet, 122-character) {
				return false
			}
		}
	}
	return true
}

func OperateSet(set map[int]struct{}, key int) bool {
	if _, ok := set[key]; ok {
		return true
	} else {
		set[key] = struct{}{}
	}
	return false
}
