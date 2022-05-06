package hamming

import "errors"

func Distance(a, b string) (distance int, err error) {
	if len(a) != len(b) {
		return 0, errors.New("Strings are of different lengths")
	}
	// In a real situation we would probably start asking about multi-byte characters.
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}
	return
}
