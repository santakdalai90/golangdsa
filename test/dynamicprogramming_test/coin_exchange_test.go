package dynamicprogramming_test

import (
	"testing"

	dynamicprogramming "github.com/santakdalai90/golangdsa/src/dynamicprogramming"

	"github.com/stretchr/testify/assert"
)

func ExampleCoinExchange() {
	dynamicprogramming.CoinExchange([]int{1, 7, 10}, 14)
	dynamicprogramming.CoinExchange([]int{1, 7, 10}, 13)
	dynamicprogramming.CoinExchange([]int{1, 7, 10}, 15)
	// Output:
	// Minimum number of coins required = 2
	// Here are the coins:
	// 7
	// 7
	// Minimum number of coins required = 4
	// Here are the coins:
	// 10
	// 1
	// 1
	// 1
	// Minimum number of coins required = 3
	// Here are the coins:
	// 7
	// 7
	// 1
}

func TestCoinExchangeNoBackTrack(t *testing.T) {
	assert.Equal(t, 2, dynamicprogramming.CoinExchangeNoBackTrack([]int{1, 7, 10}, 14))
	assert.Equal(t, 4, dynamicprogramming.CoinExchangeNoBackTrack([]int{1, 7, 10}, 13))
	assert.Equal(t, 3, dynamicprogramming.CoinExchangeNoBackTrack([]int{1, 7, 10}, 15))
	assert.Equal(t, 3, dynamicprogramming.CoinExchangeNoBackTrack([]int{2, 3, 5, 7, 8}, 18))
}

func TestCountOfWaysToExchange(t *testing.T) {
	tests := []struct {
		name          string
		denominations []int
		amount        int
		ways          int
	}{
		{"simple", []int{1, 2, 3}, 5, 5},
		{"big amount small count", []int{1, 5, 10}, 7, 2},
		{"zero value", []int{1, 2, 3}, 0, 1},
		{"big amount big count", []int{1, 2, 5}, 10, 10},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ways := dynamicprogramming.CountOfWaysToExchange(test.denominations, test.amount)
			assert.Equal(t, test.ways, ways)
		})
	}
}
