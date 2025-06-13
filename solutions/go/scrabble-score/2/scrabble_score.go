package scrabble

var magicMap map[byte]int = map[byte]int{
	65: 1, 97: 1,
	66: 3, 98: 3,
	67: 3, 99: 3,
	68: 2, 100: 2,
	69: 1, 101: 1,
	70: 4, 102: 4,
	71: 2, 103: 2,
	72: 4, 104: 4,
	73: 1, 105: 1,
	74: 8, 106: 8,
	75: 5, 107: 5,
	76: 1, 108: 1,
	77: 3, 109: 3,
	78: 1, 110: 1,
	79: 1, 111: 1,
	80: 3, 112: 3,
	81: 10, 113: 10,
	82: 1, 114: 1,
	83: 1, 115: 1,
	84: 1, 116: 1,
	85: 1, 117: 1,
	86: 4, 118: 4,
	87: 4, 119: 4,
	88: 8, 120: 8,
	89: 4, 121: 4,
	90: 10, 122: 10,
}

func Score(word string) int {
	var score int = 0
	var length int = len(word)

	for i := 0; i < length; i++ {
		score += magicMap[word[i]]
	}
	return score
}
