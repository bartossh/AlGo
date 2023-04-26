package hamming

import "errors"

// Distance calculates Hamming distance
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("DNA didn't match ")
	}
	distance := 0
	for i := 0; i < len(a); i++ {
		if byte(a[i]) != byte(b[i]) {
			distance++
		}
	}
	return distance, nil
}
