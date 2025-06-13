// Package twofer implements a solution to Two-fer exercise
package twofer

import (
	"fmt"
)

// ShareWith function accepts string parameter name and returns answer to
// exercise in string format, if name parameter is not set value "you" is
// used as default
func ShareWith(name string) string {
	if len(name) == 0 {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
