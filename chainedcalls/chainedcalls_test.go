package chainedcalls

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChainedCalls(t *testing.T) {
	var v1 int
	v2 := 10
	Lock[int]{}.Unlock().Write(v2).Read(&v1).Lock()
	assert.Equal(t, v1, v2)
}
