package util

import "fmt"

func Print2DArray[T any](arr [][]T) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Print(arr[i][j], "\t\t")
		}
		fmt.Println()
	}
}

func Create2DSlice[T any](m, n int) [][]T {
	slice2D := make([][]T, m)
	for i := range slice2D {
		slice2D[i] = make([]T, n)
	}
	return slice2D
}
