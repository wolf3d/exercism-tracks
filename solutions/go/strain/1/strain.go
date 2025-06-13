package strain

//import "fmt"

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(filter func(int) bool) Ints {
	var r Ints
    for _, k := range i {
        if filter(k) {
            r = append(r, k)
        }
    }
	return r
}

func (i Ints) Discard(filter func(int) bool) Ints {
	var r Ints
    for _, k := range i {
        if !filter(k) {
            r = append(r, k)
        }
    }
	return r
}

func (l Lists) Keep(filter func([]int) bool) Lists {
	var r Lists
    for _, k := range l {
        if filter(k) {
            r = append(r, k)
        }
    }
	return r
}

func (s Strings) Keep(filter func(string) bool) Strings {
	var r Strings
    for _, k := range s {
        if filter(k) {
            r = append(r, k)
        }
    }
	return r
}
