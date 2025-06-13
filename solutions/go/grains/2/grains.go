package grains

import (
	"errors"
)

func Square(number int) (uint64, error) {
	if number == 1 {
		return 1, nil
	}
	if number < 1 || number > 64 {
		return 0, errors.New("out of range error")
	}
	return 1 << (number - 1), nil
}

func Total() uint64 {
	return (1 << 64) - 1
}
