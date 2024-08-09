package dynamicprogramming

import "golang-dsa/src/util"

// Given two strings source and target, find the edit distance between the two.
// Permitted operations are insertion, deletion, and replacement

func EditDistance(source, target string) int {
	//make unicode compatible
	s := []rune(source) 
	t := []rune(target)

	// get lengths
	m := len(s)
	n := len(t)
	
	// create table
	E := util.Create2DSlice[int](m+1, n+1)

	// base case
	for i := 0; i <= m; i++ {
		E[i][0] = i
	}
	for j := 0; j <= n; j++ {
		E[0][j] = j
	}

	// fill table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			cost := 0
			if s[i-1] != t[j-1] {
				cost = 1
			}

			E[i][j] = util.Min(1+E[i-1][j], 1+E[i][j-1], cost+E[i-1][j-1])
		}
	}

	// solution
	return E[m][n]
}
