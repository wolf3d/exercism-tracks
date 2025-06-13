package pangram

var magicMap map[byte]byte = map[byte]byte{
	65: 97, 97: 97,
	66: 98, 98: 98,
	67: 99, 99: 99,
	68: 100, 100: 100,
	69: 101, 101: 101,
	70: 102, 102: 102,
	71: 103, 103: 103,
	72: 104, 104: 104,
	73: 105, 105: 105,
	74: 106, 106: 106,
	75: 107, 107: 107,
	76: 108, 108: 108,
	77: 109, 109: 109,
	78: 110, 110: 110,
	79: 111, 111: 111,
	80: 112, 112: 112,
	81: 113, 113: 113,
	82: 114, 114: 114,
	83: 115, 115: 115,
	84: 116, 116: 116,
	85: 117, 117: 117,
	86: 118, 118: 118,
	87: 119, 119: 119,
	88: 120, 120: 120,
	89: 121, 121: 121,
	90: 122, 122: 122,
}

//BenchmarkPangram-8         44968             25745 ns/op            1397 B/op         31 allocs/op
func IsPangram(input string) bool {
	const alphabet int = 26
	mySet := make(map[byte]struct{}, alphabet)
	length := len(input)
	if length < 26 {
		return false
	}

	for i := 0; i < length; i++ {
		var rChar byte = input[i]
		if len(mySet) == alphabet {
			return true
		}
		if rChar >= 65 && rChar <= 90 || rChar >= 97 && rChar <= 122 {
			char := magicMap[rChar]
			if _, ok := mySet[char]; !ok {
				mySet[char] = struct{}{}
			}
		}
	}
	return len(mySet) == alphabet
}
