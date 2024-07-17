package dynamicprogramming

import (
	dynamicprogramming "golang-dsa/src/dynamic_programming"
	"testing"

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
}
