// Package triangle contains methods for finding out triangle kind
package triangle

type Kind int

const (
    NaT = iota// not a triangle
    Equ // equilateral
    Iso // isosceles
    Sca // scalene
)

// KindFromSides finds out the kind of a triangle
func KindFromSides(a, b, c float64) Kind {
	var k Kind
    switch {        
    case a == 0 || b == 0 || c == 0:
        k = NaT
    case a + b < c || a + c < b || b + c < a:
    	k = NaT
    case a == b && b == c:
    	k = Equ
    case a != b && b != c && a != c:
    	k = Sca
    case 2*a >= b + c || 2*b >= a + c || 2*c >= b + a:
    	k = Iso
    }
	return k
}
