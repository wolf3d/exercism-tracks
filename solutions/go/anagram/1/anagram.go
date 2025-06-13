package anagram

import "reflect"

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

type histMap map[byte]int

func Detect(subject string, candidates []string) []string {
    subjectLowerCase := ToLowerCase(subject)
    theMap := MakeHistogram(subjectLowerCase)    
    var result []string
    for _, candidate := range candidates {
        candidateLowerCase := ToLowerCase(candidate)
        if subjectLowerCase == candidateLowerCase {
            continue
        }
        if len(subjectLowerCase) != len(candidate) {
            continue
        }
		comparison := MakeHistogram(candidateLowerCase)
        if reflect.DeepEqual(theMap, comparison) {
            result = append(result, candidate)
        }
    }
	return result
}

func MakeHistogram(subject string) histMap {
    length := len(subject)
    hsMap := make(histMap)
    for i := 0; i < length; i++ {
        char := subject[i]
		if _, ok := hsMap[char]; !ok {
            hsMap[char] = 1
        } else {
        	hsMap[char] += 1
        }
    }
	return hsMap
}

func ToLowerCase(word string) string {
    length := len(word)
    var bytes []byte
    for i := 0; i < length; i++ {
        bytes = append(bytes, magicMap[word[i]])
    }
	return string(bytes)
}
