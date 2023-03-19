package queenattack

import (
	"errors"
	"math"
	"strconv"
)

var grid = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'}

type position struct {
	x, y int
}

// CanQueenAttack returns true if queens can attack eachother or false otherwise
func CanQueenAttack(qa, qb string) (bool, error) {
	qAp, err := mapStringToPosition(qa)
	if err != nil {
		return false, err
	}
	qBp, err := mapStringToPosition(qb)
	if err != nil {
		return false, err
	}
	return qAp.attack(qBp)
}

func (a *position) attack(b *position) (bool, error) {
	if a.x == b.x && a.y == b.y {
		return false, errors.New("same square")
	}
	if a.x == b.x {
		return true, nil
	}
	if a.y == b.y {
		return true, nil
	}
	xDif := int(math.Abs(float64(a.x - b.x)))
	yDif := int(math.Abs(float64(a.y - b.y)))
	if xDif == yDif {
		return true, nil
	}
	return false, nil

}

func mapStringToPosition(s string) (*position, error) {
	if len(s) != 2 {
		return nil, errors.New("wrong position")
	}
	var x int
	found := false
	for i, v := range grid {
		if rune(s[0]) == v {
			x = i + 1
			found = true
			break
		}
	}
	if !found {
		return nil, errors.New("of board")
	}
	y, err := strconv.Atoi(string(s[1]))
	if err != nil {
		return nil, err
	}
	if y > 8 || y < 1 {
		return nil, errors.New("of board")
	}
	return &position{x, y}, nil
}
