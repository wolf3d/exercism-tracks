package atbash

import (
	"bytes"
	"strings"
)

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

//BenchmarkAtbash-8         440127              2765 ns/op            1344 B/op         40 allocs/op

const magicConst byte = 122

func Atbash(s string) string {
	var space []byte = make([]byte, 1)
	space[0] = ' '

	lengthS := len(s)
	var newBytes bytes.Buffer
	builder := strings.Builder{}

	for i := 0; i < lengthS; i++ {
		char := s[i]
		if char >= 97 && char <= 122 {
			newBytes.WriteByte(magicConst - (char - 'a'))
		}
		if char >= 65 && char <= 90 {
			char = magicMap[char]
			newBytes.WriteByte(magicConst - (char - 'a'))
		}
		if char >= 48 && char <= 57 {
			newBytes.WriteByte(char)
		}
	}

	slist := newBytes.Next(5)
	for lengthBytes := len(slist); lengthBytes > 0; {
		builder.Write(slist)
		slist = newBytes.Next(5)
		if lengthBytes == 5 && len(slist) > 0 {
			builder.Write(space)
		}
		lengthBytes = len(slist)
	}
	return builder.String()
}
