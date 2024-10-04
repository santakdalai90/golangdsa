package dynamicprogramming

import (
	"math"

	"golangdsa/util"
)

// Given two strings find the longest common subsequence and its length

func LCS(X, Y string) (int, string) {
	m, n := len(X), len(Y)
	table := util.Create2DSlice[int](m+1, n+1)

	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if i == 0 || j == 0 {
				table[i][j] = 0
			} else if X[i-1] == Y[j-1] {
				table[i][j] = 1 + table[i-1][j-1]
			} else {
				table[i][j] = int(math.Max(float64(table[i][j-1]), float64(table[i-1][j])))
			}
		}
	}

	l := ""
	for i, j := m, n; i > 0 && j > 0; {
		// fmt.Printf("Location: %d,%d\n", i, j)
		if X[i-1] == Y[j-1] {
			// fmt.Println("picked:", string(X[i-1]))
			l = string(X[i-1]) + l
			i--
			j--
		} else {
			if table[i-1][j] > table[i][j-1] {
				i--
			} else {
				j--
			}
		}
	}

	return table[m][n], l
}

func LCSOptimized(X, Y string) int {
	m, n := len(X), len(Y)
	table := util.Create2DSlice[int](2, n+1)
	// base case
	for j := 0; j <= n; j++ {
		table[0][j] = 0

	}
	table[1][0] = 0
	// fill table
	for i := 1; i <= m; i++ {
		row := i % 2
		for j := 1; j <= n; j++ {
			if X[i-1] == Y[j-1] {
				table[row][j] = 1 + table[util.Abs(row-1)][j-1]
			} else {
				table[row][j] = int(math.Max(float64(table[row][j-1]), float64(table[util.Abs(row-1)][j])))
			}
		}
	}
	return table[m%2][n]
}
