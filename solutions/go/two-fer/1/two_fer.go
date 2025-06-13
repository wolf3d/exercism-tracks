/*
package twofer implements a solution to Two-fer exercise

Introduction

Two-fer or 2-fer is short for two for one. One for you
 and one for me.

"One for X, one for me."

When X is a name or "you".

If the given name is "Alice", the result should be
"One for Alice, one for me." If no name is given,
the result should be "One for you, one for me."

*/
package twofer

import (
	"fmt"
)

// ShareWith string parameter name, returns result string. Checks if name
// parameter is empty, assigns default value "you" if name is empty.
// Sprintf formats string "One for X, one for me." according to a format
// specifier and then resulting string is returned
func ShareWith(name string) string {
	if len(name) == 0 {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
