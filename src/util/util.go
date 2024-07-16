package util

import "fmt"

func Print2DArray(arr [][]int) {
	for i:=0; i<len(arr); i++ {
		for j:=0; j<len(arr[i]); j++ {
			fmt.Print(arr[i][j], "\t\t")
		}
		fmt.Println()
	}
}