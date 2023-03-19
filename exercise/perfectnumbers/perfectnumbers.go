package perfect

import "fmt"

// Classification describes number by Nicomachus perfect number categories
type Classification int

var (
	// ClassificationPerfect sum of factors of a number is equal to number
	ClassificationPerfect Classification = 0
	// ClassificationDeficient sum of factors of a number is smaller to number
	ClassificationDeficient Classification = -1
	// ClassificationAbundant sum of factors of a number is bigger to number
	ClassificationAbundant Classification = 1
)

// ErrOnlyPositive only positive numbers description
var ErrOnlyPositive error = fmt.Errorf("only positive numbers are allowed")

func factors(n int64) []int64 {
	res := []int64{}
	for t := n - 1; t > 0; t-- {
		if n%t == 0 {
			res = append(res, t)
		}
	}
	return res
}

func sumFactors(f []int64) (sum int64) {
	for _, v := range f {
		sum += v
	}
	return sum
}

// Classify classifies the number according to Nicomachus perfect number categories
func Classify(n int64) (c Classification, err error) {
	if n < 1 {
		return 0, ErrOnlyPositive
	}
	sum := sumFactors(factors(n))
	if sum == n {
		return ClassificationPerfect, nil
	}
	if sum > n {
		return ClassificationAbundant, nil
	}
	return ClassificationDeficient, nil
}
