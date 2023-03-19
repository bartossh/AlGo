package palindrome

import (
	"fmt"
)

// Product holds information about palindrom
type Product struct {
	Product        int      // palindromic, of course
	Factorizations [][2]int //list of all possible two-factor factorizations of Product, within given limits, in order
}

func isPalindrom(n int) bool {
	if n < 10 {
		return true
	}
	var reverse, remainder, number int = 0, 0, n
	for {
		remainder = number % 10
		reverse = reverse*10 + remainder
		number /= 10

		if number == 0 {
			break
		}
	}
	if reverse == n {
		return true
	}
	return false
}

func addToPalindroms(p *map[int][][2]int, n, j, i int) {
	if v, ok := (*p)[n]; ok {
		(*p)[n] = append(v, [2]int{j, i})
	} else {
		(*p)[n] = [][2]int{{j, i}}
	}
}

// Products calculates min and max Palindrom
func Products(fmin, fmax int) (pmin, pmax Product, err error) {
	if fmin > fmax {
		return pmin, pmax, fmt.Errorf("fmin > fmax")
	}
	palindroms := make(map[int][][2]int)
	j := fmin
	for j <= fmax {
		for i := j; i <= fmax; i++ {
			product := j * i
			if isPalindrom(product) {
				addToPalindroms(&palindroms, product, j, i)
			}
		}
		j++
	}
	for k, v := range palindroms {
		if fmin > 1 {
			if pmin.Product == 0 {
				pmin.Product = k
				pmin.Factorizations = v
			} else if pmin.Product >= k {
				pmin.Product = k
				pmin.Factorizations = v
			}
		}
		if pmax.Product == 0 {
			pmax.Product = k
			pmax.Factorizations = v
		} else if pmax.Product <= k {
			pmax.Product = k
			pmax.Factorizations = v
		}
	}
	if len(pmax.Factorizations) == 0 {
		return pmin, pmax, fmt.Errorf("no palindromes")
	}
	return
}
