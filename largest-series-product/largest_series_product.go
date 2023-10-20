package lsproduct

import (
	"errors"
	"strconv"
)

// LargestSeriesProduct calculates the largest series product for givendigits.
func LargestSeriesProduct(digits string, span int) (int64, error) {
	if span < 0 {
		return 0, errors.New("span must not be negative")
	}
	if span > len(digits) {
		return 0, errors.New("span must be smaller then digits length")
	}

	var product int64
	var start int
	for start <= len(digits)-span {
		series := digits[start : start+span]
		var local int64 = 1
		for _, r := range series {
			digit, err := strconv.Atoi(string(r))
			if err != nil {
				return 0, errors.New("digits input must only contain digits")
			}
			local *= int64(digit)
		}
		if local > product {
			product = local
		}
		start++
	}
	return product, nil
}
