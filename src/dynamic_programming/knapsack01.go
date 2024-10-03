package dynamicprogramming

import (
	"fmt"
	"github.com/santakdalai90/golang-dsa/src/util"
)

// Given a knapsack with a capacity and some items with given weight and value,
// write a code to maximize the value of the knapsack. You can take an item or not, but not in fractions.
// Also print the chosen items.

type Item struct {
	Weight int
	Value  int
}

func Knapsack01(knapsackSize int, items []Item) {
	n := len(items)
	// create table
	table := util.Create2DSlice[int](n+1, knapsackSize+1)

	// fill table
	for i := 1; i <= n; i++ {
		for j := 1; j <= knapsackSize; j++ {
			if j >= items[i-1].Weight && items[i-1].Value+table[i-1][j-items[i-1].Weight] > table[i-1][j] {
				table[i][j] = items[i-1].Value + table[i-1][j-items[i-1].Weight]
			} else {
				table[i][j] = table[i-1][j]
			}
		}
	}

	// print the max value of knapsack
	fmt.Println("Maximum possible value of the knapsack:", table[n][knapsackSize])

	// backtracking
	for i, j := n, knapsackSize; i > 0 && j > 0; i-- {
		if table[i][j] != table[i-1][j] {
			fmt.Printf("Selected item no. %d of weight: %d and value: %d\n", i, items[i-1].Weight, items[i-1].Value)
			j -= items[i-1].Weight
		}
	}

}
