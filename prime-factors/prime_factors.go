package prime

type primes struct {
	inner    []int64
	position int
}

func newPrimes() *primes {
	inner := make([]int64, 1, 10)
	inner[0] = 2
	return &primes{inner, 0}
}

func (p *primes) nextPrime(n int64) (int64, bool) {
	if p.position < len(p.inner) {
		v := p.inner[p.position]
		p.position++
		return v, true
	}
	start := p.inner[len(p.inner)-1] + 1
	for candidate := start; candidate <= n; candidate++ {
		isPrime := true
	Check:
		for _, v := range p.inner {
			if candidate%v == 0 {
				isPrime = false
				break Check
			}
		}
		if isPrime {
			p.inner = append(p.inner, candidate)
			p.position = len(p.inner) - 1
			return candidate, true
		}
	}
	return 0, false
}

func (p *primes) resetPosition() {
	p.position = 0
}

func Factors(n int64) []int64 {
	factors := make([]int64, 0)
	p := newPrimes()
	for {
		v, ok := p.nextPrime(n)
		if !ok {
			break
		}
		if n%v == 0 {
			factors = append(factors, v)
			n = n / v
			p.resetPosition()
			continue
		}
	}
	return factors
}
