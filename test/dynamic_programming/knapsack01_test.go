package dynamicprogramming

import (
	dynamicprogramming "github.com/santakdalai90/GolangDSA/src/dynamic_programming"
)

func ExampleKnapsack01() {
	items := []dynamicprogramming.Item{
		{Weight: 2, Value: 12},
		{Weight: 1, Value: 10},
		{Weight: 3, Value: 21},
		{Weight: 2, Value: 15},
	}
	dynamicprogramming.Knapsack01(7, items)

	items = []dynamicprogramming.Item{
		{Weight: 4, Value: 12},
		{Weight: 6, Value: 10},
		{Weight: 5, Value: 8},
		{Weight: 7, Value: 11},
		{Weight: 3, Value: 14},
		{Weight: 1, Value: 7},
		{Weight: 6, Value: 9},
	}
	dynamicprogramming.Knapsack01(18, items)

	items = []dynamicprogramming.Item{
		{Weight: 5, Value: 12},
		{Weight: 7, Value: 10},
		{Weight: 3, Value: 9},
		{Weight: 8, Value: 16},
		{Weight: 4, Value: 14},
		{Weight: 3, Value: 7},
		{Weight: 7, Value: 14},
		{Weight: 5, Value: 10},
	}
	dynamicprogramming.Knapsack01(23, items)

	// Output:
	// Maximum possible value of the knapsack: 48
	// Selected item no. 4 of weight: 2 and value: 15
	// Selected item no. 3 of weight: 3 and value: 21
	// Selected item no. 1 of weight: 2 and value: 12
	// Maximum possible value of the knapsack: 44
	// Selected item no. 5 of weight: 3 and value: 14
	// Selected item no. 3 of weight: 5 and value: 8
	// Selected item no. 2 of weight: 6 and value: 10
	// Selected item no. 1 of weight: 4 and value: 12
	// Maximum possible value of the knapsack: 58
	// Selected item no. 6 of weight: 3 and value: 7
	// Selected item no. 5 of weight: 4 and value: 14
	// Selected item no. 4 of weight: 8 and value: 16
	// Selected item no. 3 of weight: 3 and value: 9
	// Selected item no. 1 of weight: 5 and value: 12
}
