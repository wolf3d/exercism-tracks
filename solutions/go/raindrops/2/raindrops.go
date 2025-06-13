package raindrops

import (
	"fmt"
	"strings"
)

func Convert(number int) string {
	var sb strings.Builder

	if number%3 == 0 {
		sb.WriteString("Pling")
	}
	if number%5 == 0 {
		sb.WriteString("Plang")
	}
	if number%7 == 0 {
		sb.WriteString("Plong")
	}
	if sb.Len() == 0 {
		return fmt.Sprint(number)
	}
	return sb.String()
}
