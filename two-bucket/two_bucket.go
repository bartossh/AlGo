package twobucket

import (
	"errors"
)

type bucket struct {
	size, amount int
	name         string
}

func Solve(sizeOne, sizeTwo, goal int, startBucket string) (goalBucket string, moves int, otherBucket int, err error) {
	if sizeOne < 1 || sizeTwo < 1 {
		err = errors.New("bad bucket")
		return
	}
	for i := 2; i <= sizeOne && i <= sizeTwo; i++ {
		if sizeOne%i == 0 && sizeTwo%i == 0 {
			err = errors.New("bucket sizes are not relative prime")
			return
		}
	}
	if goal < 1 {
		err = errors.New("bad goal")
		return
	}
	if startBucket != "one" && startBucket != "two" {
		err = errors.New("invalid bucket name")
		return
	}

	var a, b *bucket
	if startBucket == "one" {
		a, b = &bucket{size: sizeOne, amount: sizeOne, name: "one"}, &bucket{size: sizeTwo, amount: 0, name: "two"}
	} else {
		a, b = &bucket{size: sizeOne, amount: 0, name: "one"}, &bucket{size: sizeTwo, amount: sizeTwo, name: "two"}
		a, b = b, a
	}
	moves++
	for {
		if a.amount == goal {
			goalBucket = a.name
			otherBucket = b.amount
			break
		}
		if b.amount == goal {
			goalBucket = b.name
			otherBucket = a.amount
			break
		}
		if transfer(a, b, goal) {
			a, b = b, a
		}
		moves++
	}
	return
}

func transfer(a, b *bucket, goal int) bool {
	if b.amount < b.size && b.size == goal {
		b.amount = b.size
		return false
	}

	if a.amount < a.size && a.size == goal {
		a.amount = a.size
		return false
	}

	if a.amount == 0 {
		a.amount = a.size
		return false
	}

	if a.size > b.size && b.amount == b.size {
		b.amount = 0
		return false
	}

	if b.size > a.size && b.amount == b.size {
		b.amount = 0
		return false
	}

	if a.amount == a.size && b.amount == b.size {
		a.amount = 0
		return false
	}

	delta := b.size - b.amount
	if delta >= a.amount {
		b.amount += a.amount
		a.amount = 0
	} else {
		b.amount += delta
		a.amount -= delta
	}
	return a.amount == a.size
}
