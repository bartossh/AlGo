// Package react implements a basic reactive system.
package react

import (
	"crypto/rand"
	"encoding/binary"
)

type (
	reactor struct {
		registry []*computeCell
	}
	inputCell struct {
		value int
		r     *reactor
	}
	computeCell struct {
		compute   func() int
		callbacks map[int]func(int)
		prev      int
	}

	canceler struct {
		id           int
		registeredTo *computeCell
	}
)

func generateID() int {
	token := make([]byte, 8)
	rand.Read(token)
	return int(binary.BigEndian.Uint64(token))
}

func (ic *inputCell) SetValue(i int) {
	ic.value = i
	for _, cc := range ic.r.registry {
		if len(cc.callbacks) == 0 || cc.Value() == cc.prev {
			continue
		}
		for _, cb := range cc.callbacks {
			cb(cc.Value())
		}
		cc.prev = cc.Value()
	}
}

func (ic *inputCell) Value() int {
	return ic.value
}

func (cc *computeCell) Value() int {
	return cc.compute()
}

func (cc *computeCell) AddCallback(cb func(int)) Canceler {
	id := generateID()
	cc.callbacks[id] = cb
	return &canceler{
		id:           id,
		registeredTo: cc,
	}
}

func (c *canceler) Cancel() {
	delete(c.registeredTo.callbacks, c.id)
}

func New() Reactor {
	return &reactor{
		registry: make([]*computeCell, 0),
	}
}

func (r *reactor) CreateInput(i int) InputCell {
	return &inputCell{
		value: i,
		r:     r,
	}
}

func (r *reactor) CreateCompute1(c Cell, compute func(int) int) ComputeCell {
	return r.createCompute(
		func() int {
			return compute(c.Value())
		},
	)
}

func (r *reactor) CreateCompute2(c1, c2 Cell, compute func(int, int) int) ComputeCell {
	return r.createCompute(
		func() int {
			return compute(c1.Value(), c2.Value())
		},
	)
}
func (r *reactor) createCompute(compute func() int) *computeCell {
	var cc computeCell
	cc.compute = compute
	cc.callbacks = make(map[int]func(int))
	cc.prev = cc.Value()
	r.registry = append(r.registry, &cc)
	return &cc
}
