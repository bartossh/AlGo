package coinchange

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuccess(t *testing.T) {

	cases := []struct {
		coins         []uint
		amount, value uint
	}{
		{
			coins:  []uint{1, 2, 5, 10},
			amount: 101,
			value:  11,
		},
		{
			coins:  []uint{1, 2, 5, 10},
			amount: 55,
			value:  6,
		},
		{
			coins:  []uint{1, 2, 5, 10},
			amount: 12,
			value:  2,
		},
		{
			coins:  []uint{1, 2, 5, 10},
			amount: 22,
			value:  3,
		},
		{
			coins:  []uint{1, 2, 5, 10},
			amount: 123,
			value:  14,
		},
		{
			coins:  []uint{1, 2, 5, 10, 20, 50, 100, 200, 500},
			amount: 1203,
			value:  5,
		},
		{
			coins:  []uint{1, 2, 5, 10, 20, 50, 100, 200, 500},
			amount: 1234,
			value:  7,
		},
		{
			coins:  []uint{1, 2, 5, 10, 20, 50, 100, 200, 500},
			amount: 5555,
			value:  13,
		},
		{
			coins:  []uint{200, 500},
			amount: 2000,
			value:  4,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("testing coins %v for amount %v", c.coins, c.amount), func(t *testing.T) {
			v, err := CoinChange(c.coins, c.amount)
			assert.Equal(t, c.value, v)
			assert.Nil(t, err)
		})
	}
}

func TestFailure(t *testing.T) {

	cases := []struct {
		coins  []uint
		amount uint
	}{
		{
			coins:  []uint{5, 10},
			amount: 103,
		},
		{
			coins:  []uint{200, 500},
			amount: 2001,
		},
		{
			coins:  []uint{200, 500},
			amount: 10,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("testing coins %v for amount %v", c.coins, c.amount), func(t *testing.T) {
			_, err := CoinChange(c.coins, c.amount)
			assert.NotNil(t, err)
		})
	}
}

func BenchmarkCoinChangeLarge(b *testing.B) {
	coins := []uint{1, 2, 5, 10, 20, 50, 100, 200, 500}
	var amount uint = 5555

	for n := 0; n < b.N; n++ {
		CoinChange(coins, amount)
	}
}
