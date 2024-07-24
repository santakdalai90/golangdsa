package util

import "fmt"

func Print2DArray(arr [][]int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Print(arr[i][j], "\t\t")
		}
		fmt.Println()
	}
}

func Create2DSlice(m, n int) [][]int {
	slice2D := make([][]int, m)
	for i := range slice2D {
		slice2D[i] = make([]int, n)
	}
	return slice2D
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
