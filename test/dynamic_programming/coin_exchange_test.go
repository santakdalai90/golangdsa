package dynamicprogramming

import (
	dynamicprogramming "golang-dsa/src/dynamic_programming"
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
