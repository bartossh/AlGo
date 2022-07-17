package coinchange

import (
	"fmt"
)

const maxUint = ^uint(0)

// CoinChange calculates minimal amount of coins to be used to get amount if possible or returns error otherwise
func CoinChange(coins []uint, amount uint) (uint, error) {
	mem := make(map[uint]uint, amount+1)
	mem[0] = 0

	var i uint = 1
	for i <= amount {
		mem[i] = maxUint
		for _, c := range coins {
			if c <= i && mem[i-c] != maxUint {
				mem[i] = min(mem[i], mem[i-c]+1)
			}
		}
		i++
	}
	if mem[amount] == maxUint {
		return 0, fmt.Errorf("amount %v cannot be constructed from provided coins", amount)
	}
	return mem[amount], nil
}

func min(a, b uint) uint {
	if a < b {
		return a
	}
	return b
}
