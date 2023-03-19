package diffiehellman

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// PrivateKey generates private key
func PrivateKey(p *big.Int) *big.Int {
	if p.Sign() != 1 {
		panic(fmt.Sprintf("error when generating big int random number, p value equal to or less then 0"))
	}
	m := &big.Int{}
	m.Sub(p, big.NewInt(2))
	n, err := rand.Int(rand.Reader, m)
	if err != nil {
		panic(fmt.Sprintf("error when generating big int random number: %v", err))
	}
	r := &big.Int{}
	r.Add(n, big.NewInt(2))
	return r
}

// PublicKey returns calculated public key: A = g**a mod p
func PublicKey(private, p *big.Int, g int64) *big.Int {
	gi := big.NewInt(g)
	e := &big.Int{}
	e.Exp(gi, private, p)
	return e
}

// NewPair creates pair of private and public keys
func NewPair(p *big.Int, g int64) (private, public *big.Int) {
	private = PrivateKey(p)
	public = PublicKey(private, p, g)
	return
}

// SecretKey creates secret key: s = B**a mod p
func SecretKey(private1, public2, p *big.Int) *big.Int {
	s := &big.Int{}
	s.Exp(public2, private1, p)
	return s
}
