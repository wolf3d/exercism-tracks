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

func IsPangram(input string) bool {
	mySet := make(map[byte]struct{})
	length := len(input)
    if length < 26 {
        return false
    }
	const alphabet int = 26    
	for i := 0; i < length; i++ {        
        if input[i] == 20 {
            continue
        }
        if char, ok := magicMap[input[i]]; ok {
	        if _, ok := mySet[char]; !ok {
	        	mySet[char] = struct{}{} 
    	    }            
        }
	}
	return len(mySet) == alphabet
}
