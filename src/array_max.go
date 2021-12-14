package src

import (
	"errors"
)

// ArrayMax finds the maximum element in an array
// if array is blank or nil then returns -1, error
func ArrayMax(arr []int) (int, error) {
	if len(arr) == 0 {
		return -1, errors.New(ErrorEmptyArray)
	}

	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max, nil
}
