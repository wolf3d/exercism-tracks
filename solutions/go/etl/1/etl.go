package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	var out map[string]int = make(map[string]int)
    for k, v := range in {
        for _, i := range v {
            out[strings.ToLower(i)] = k
        }
    }
    return out
}
