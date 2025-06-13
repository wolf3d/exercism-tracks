package grains

import "errors"

func Square(number int) (uint64, error) {
    if number == 1 {
        return uint64(number), nil
    }
if number < 1 || number > 64 {
    return 0, errors.New("out of range error")
}
	return 1 << (number-1), nil
}

func Total() uint64 {
    var result uint64 = 1
    for i := 1; i < 65; i++ {
        result += 1 << i
    }
	return result
}
