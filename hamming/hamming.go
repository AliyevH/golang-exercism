package hamming

import (
	"errors"
)

// Distance function finds hamming distance
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("Mismatch of DNA lengths")
	}
	hamming := 0
	for i := 0; i <= len(a)-1; i++ {
		if a[i] != b[i] {
			hamming++
		}
	}
	return hamming, nil
}
