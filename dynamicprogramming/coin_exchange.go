package dynamicprogramming

import (
	"fmt"

	"github.com/santakdalai90/golangdsa/util"
)

// Given an amount and you have infinite number of coins of given denominations, you have to find the minimum
// number of coins which sums up to the amount and also print the coins involved.

func CoinExchange(denominations []int, amount int) {
	// initialize table
	table := util.Create2DSlice[int](amount+1, len(denominations))

	for i := 0; i < amount+1; i++ {
		table[i][0] = i
	}

	for j := 0; j < len(denominations); j++ {
		table[0][j] = 0
	}

	// fill the table with minimum coins
	for i := 1; i <= amount; i++ {
		for j := 1; j < len(denominations); j++ {
			if i >= denominations[j] && table[i-denominations[j]][j]+1 < table[i][j-1] {
				table[i][j] = table[i-denominations[j]][j] + 1
			} else {
				table[i][j] = table[i][j-1]
			}
		}
	}

	// util.Print2DArray(table)

	fmt.Println("Minimum number of coins required =", table[amount][len(denominations)-1])

	// backtrack to get the coins
	fmt.Println("Here are the coins:")
	i := amount
	j := len(denominations) - 1
	for i > 0 && j > 0 {
		if table[i][j] == table[i][j-1] {
			j--
		} else {
			fmt.Println(denominations[j])
			i -= denominations[j]
		}
	}

	for i > 0 {
		fmt.Println(denominations[0])
		i--
	}
}

// CoinExchangeNoBackTrack returns the minimum number of coins required without giving the denominations
func CoinExchangeNoBackTrack(denominations []int, amount int) int {
	// create 1-dimensional array

	table := make([]int, amount+1)

	j := 0
	for i := 0; i <= amount; i++ {
		table[i] = i
	}

	for j = 1; j < len(denominations); j++ {
		for i := 1; i <= amount; i++ {
			if (i >= denominations[j]) && (table[i-denominations[j]]+1 < table[i]) {
				table[i] = table[i-denominations[j]] + 1
			}
		}
	}

	return table[amount]
}

// CountOfWaysToExchange returns the number of possible ways the given amount
// be exchanged with the availabel denominations
func CountOfWaysToExchange(denominations []int, amount int) int {
	// create table
	table := util.Create2DSlice[int](len(denominations), amount+1)

	// base case
	for i := 0; i < len(denominations); i++ {
		table[i][0] = 1
	}

	for j := 1; j <= amount; j++ {
		table[0][j] = 1
	}

	// fill the table with count of solutions
	for i := 1; i < len(denominations); i++ {

		for j := 1; j <= amount; j++ {
			if denominations[i] <= j {
				// fmt.Printf("denominations: %d, i: %d, j: %d\n", len(denominations), i, j)
				table[i][j] = table[i-1][j] + table[i][j-denominations[i]]
			} else {
				table[i][j] = table[i-1][j]
			}
		}
	}

	// util.Print2DArray(table)

	return table[len(denominations)-1][amount]
}
