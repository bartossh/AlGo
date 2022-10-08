package v8random

func XorShift128(state0, state1 *uint64) {
	s1 := *state0
	s0 := *state1
	*state0 = s0
	s1 ^= s1 << 23
	s1 ^= s1 >> 17
	s1 ^= s0
	s1 ^= s0 >> 26
	*state1 = s1
}
