package hamming

import "errors"

func Distance(a, b string) (int, error) {
	var len_a, len_b int = len(a), len(b)
    var hamming int
    if len_a != len_b {
        return 0, errors.New("strings are not equal length")        
    }
	if len_a == 0 {
        return 0, nil
    }
	for i := 0; i < len_a; i++ {
        if a[i] != b[i] {
            hamming++
        }
    }
	return hamming, nil
}
