package raindrops

import (    
    "fmt"
    "bytes"
)

func Convert(number int) string {
    var buffer bytes.Buffer
    if number % 3 == 0 {
        buffer.WriteString("Pling")
	}
	if number % 5 == 0 {
    	buffer.WriteString("Plang")
    }
	if number % 7 == 0 {
    	buffer.WriteString("Plong")
    }
	if buffer.Len() == 0 {
    	return fmt.Sprint(number)
    }
	return buffer.String()
}
