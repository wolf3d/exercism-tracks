package rotationalcipher

import "strings"

func RotationalCipher(plain string, shiftKey int) string {
	const basePosition byte = 96
    const basePositionCaps byte = 64
	const upperLimit byte = 122
    const upperLimitCaps byte = 90
    rotated := strings.Builder{}

	lenght := len(plain)
	var rot byte = byte(shiftKey % 26)    

	for i := 0; i < lenght; i++ {
        if plain[i] > basePosition && plain[i] <= upperLimit {
            rotated.WriteByte(CalculateRotation(plain[i], rot, upperLimit, basePosition))
    	} else if plain[i] > basePositionCaps && plain[i] <= upperLimitCaps {
        	rotated.WriteByte(CalculateRotation(plain[i], rot, upperLimitCaps, basePositionCaps))
		} else {
    		rotated.WriteByte(plain[i])
    	}
	}
	return rotated.String()
}

func CalculateRotation(char byte, rot byte, upperLimit byte, basePosition byte) byte {
	if char+rot > upperLimit {
		left := char + rot - upperLimit
        	return basePosition+left
	}
	return char+rot
}
