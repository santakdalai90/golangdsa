package dynamicprogramming

import (
	"golang-dsa/src/util"
)

// Given a set and a number, figure out if there exists a subset
// so that sum of its elements is equal to the given number

func IsSubsetSum(arr []int, sum int) bool{
	n := len(arr)

	// create table
	table := util.Create2DSlice[bool](n+1, sum+1)

	// base case
	for i:=0; i<=n; i++ {
		table[i][0] = true
	}
	for j:=1; j<=sum; j++ {
		table[0][j] = false
	}

	// fill table
	for i:=1; i <=n; i++ {
		for j:=1; j<=sum; j++ {
			if arr[i-1] <= j {
				table[i][j] = table[i-1][j] || table[i-1][j-arr[i-1]]
			} else {
				table[i][j] = table[i-1][j]
			}
		}
	}

	//result
	return table[n][sum]
}
