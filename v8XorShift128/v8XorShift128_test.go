package v8random

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXorShift128(t *testing.T) {
	var s0, s1 uint64 = 6636448563767554476, 11374681941261342003

	var state0, state1 uint64 = 1, 0

	for i := 0; i < 100; i++ {
		XorShift128(&state0, &state1)
		s0, s1 = state0, state1
	}
	assert.Equal(t, s0, state0)
	assert.Equal(t, s1, state1)
}
