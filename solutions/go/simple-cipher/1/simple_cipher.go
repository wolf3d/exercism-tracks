package cipher

import (
	"strings"
)

//import "fmt"

var magicMap map[byte]int = map[byte]int{
	65: 1, 97: 1,
	66: 2, 98: 2,
	67: 3, 99: 3,
	68: 4, 100: 4,
	69: 5, 101: 5,
	70: 6, 102: 6,
	71: 7, 103: 7,
	72: 8, 104: 8,
	73: 9, 105: 9,
	74: 10, 106: 10,
	75: 11, 107: 11,
	76: 12, 108: 12,
	77: 13, 109: 13,
	78: 14, 110: 14,
	79: 15, 111: 15,
	80: 16, 112: 16,
	81: 17, 113: 17,
	82: 18, 114: 18,
	83: 19, 115: 19,
	84: 20, 116: 20,
	85: 21, 117: 21,
	86: 22, 118: 22,
	87: 23, 119: 23,
	88: 24, 120: 24,
	89: 25, 121: 25,
	90: 26, 122: 26,
}

var vMap map[byte]int = map[byte]int{
	97:  1,
	98:  2,
	99:  3,
	100: 4,
	101: 5,
	102: 6,
	103: 7,
	104: 8,
	105: 9,
	106: 10,
	107: 11,
	108: 12,
	109: 13,
	110: 14,
	111: 15,
	112: 16,
	113: 17,
	114: 18,
	115: 19,
	116: 20,
	117: 21,
	118: 22,
	119: 23,
	120: 24,
	121: 25,
	122: 26,
}

const basePosition byte = 97
const basePositionCaps byte = 65
const upperLimit byte = 122
const upperLimitCaps byte = 90

const start int = 1
const end int = 26

// Define the shift and vigenere types here.
// Both types should satisfy the Cipher interface.

type shift struct {
	distance int
}

type vigenere struct {
	key string
}

func NewCaesar() Cipher {
	return shift{distance: 3}
}

func NewShift(distance int) Cipher {
	if distance >= 1 && distance <= 25 || distance <= -1 && distance >= -25 {
		return shift{distance: distance}
	} else {
		return nil
	}
}

func (c shift) Encode(input string) string {
	rotated := strings.Builder{}
	length := len(input)
	for i := 0; i < length; i++ {
		if input[i] >= basePosition && input[i] <= upperLimit || input[i] >= basePositionCaps && input[i] <= upperLimitCaps {
			src := magicMap[input[i]]
			rotated.WriteByte(byte(CalculateRotation(src, c.distance, end, start)))
		}
	}
	return rotated.String()
}

func (c shift) Decode(input string) string {
	rotated := strings.Builder{}
	length := len(input)
	for i := 0; i < length; i++ {
		if input[i] >= basePosition && input[i] <= upperLimit || input[i] >= basePositionCaps && input[i] <= upperLimitCaps {
			src := magicMap[input[i]]
			rotated.WriteByte(byte(CalculateDeRotation(src, c.distance, end, start)))
		}
	}
	return rotated.String()
}

func NewVigenere(key string) Cipher {
	length := len(key)
	fMap := make(map[byte]int)
	if length > 0 {
		for i := 0; i < length; i++ {
			_, exists := vMap[key[i]]
			if !exists {
				return nil
			}
			if length == 1 && key[i] == 97 {
				return nil
			}
			v, exists2 := fMap[key[i]]
			if length == 2 && exists2 && v == 1 {
				return nil
			}
			fMap[key[i]] = 1
		}
		return vigenere{key: key}
	}
	return nil
}

func (v vigenere) Encode(input string) string {
	rotated := strings.Builder{}
	length := len(input)
	kLength := len(v.key)
	var k int = 0
	for i := 0; i < length; i++ {
		if input[i] >= basePosition && input[i] <= upperLimit || input[i] >= basePositionCaps && input[i] <= upperLimitCaps {
			distance := vMap[v.key[k]] - 1
			if k+1 == kLength {
				k = 0
			} else {
				k += 1
			}
			src := magicMap[input[i]]
			rotated.WriteByte(byte(CalculateRotation(src, distance, end, start)))
		}
	}
	return rotated.String()
}

func (v vigenere) Decode(input string) string {
	rotated := strings.Builder{}
	length := len(input)
	kLength := len(v.key)
	var k int = 0
	for i := 0; i < length; i++ {
		if input[i] >= basePosition && input[i] <= upperLimit || input[i] >= basePositionCaps && input[i] <= upperLimitCaps {
			distance := vMap[v.key[k]] - 1
			if k+1 == kLength {
				k = 0
			} else {
				k += 1
			}
			src := magicMap[input[i]]
			rotated.WriteByte(byte(CalculateDeRotation(src, distance, end, start)))
		}
	}
	return rotated.String()
}

func CalculateRotation(char int, rot int, upperLimit int, basePosition int) int {
	if rot < 0 && char+rot < basePosition {
		return 96 + upperLimit + char + rot
	}
	if rot > 0 && char+rot > upperLimit {
		return 96 + (char+rot)%upperLimit
	}
	return 96 + char + rot
}

func CalculateDeRotation(char int, rot int, upperLimit int, basePosition int) int {
	if rot < 0 && char-rot > upperLimit {
		return 96 + (char-rot)%upperLimit
	}
	if rot > 0 && char-rot < basePosition {
		return 96 + upperLimit + char - rot
	}
	return 96 + char - rot
}
