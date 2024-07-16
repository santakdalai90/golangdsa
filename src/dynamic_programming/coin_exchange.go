package dynamicprogramming

import "fmt"

// Given an amount and you have infinite number of coins of given denominations, you have to find the minimum
// number of coins which sums up to the amount and also print the coins involved.

func CoinExchange(denominations []int, amount int) {
	// initialize table
	table := make([][]int, amount+1)
	for i := range table {
		table[i] = make([]int, len(denominations))
	}

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
