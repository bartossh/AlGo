// Based on https://en.wikipedia.org/wiki/Fermat%27s_factorization_method
// With common RSA key sizes (2048 bit) in tests,
// the Fermat algorithm with 100 rounds reliably factors numbers where p and q differ up to 2^517. I
// n other words, it can be said that primes that only differ within the lower 64 bytes
// (or around half their size) will be vulnerable.
package crackrsa

import (
	"errors"
	"math/big"
)

const iterations = 1000

// CrackRSA algorithm is cracking RSA tool when p and q are not too far apart.
// This will not crack topically used RSA keys as p and q are picked to be very far apart.
// This is a toy algorithm not a real cracking tool. If it cracks your key, you are using insecure RSA algorithm.
func CrackRSA(e, n []byte) ([]byte, error) {
	nn := (&big.Int{}).SetBytes(n)
	a := (&big.Int{}).Add((&big.Int{}).Sqrt(nn), big.NewInt(1))

	b := &big.Int{}

	for i := iterations; i > 0; i-- {
		b2 := (&big.Int{}).Mul(a, a)
		b2 = (&big.Int{}).Sub(b2, nn)
		c := (&big.Int{}).Sqrt(b2)
		if (&big.Int{}).Mul(c, c).Cmp(b2) == 0 {
			b = c
			break
		}
		a = (&big.Int{}).Add(a, big.NewInt(1))
	}

	p := (&big.Int{}).Add(a, b)
	q := (&big.Int{}).Sub(a, b)

	if (&big.Int{}).Mul(p, q).Cmp(nn) != 0 {
		return nil, errors.New("p and q do not multiply to n, cannot crack")
	}

	phi := (&big.Int{}).Mul((&big.Int{}).Sub(p, big.NewInt(1)), (&big.Int{}).Sub(q, big.NewInt(1)))
	return (&big.Int{}).ModInverse((&big.Int{}).SetBytes(e), phi).Bytes(), nil
}
