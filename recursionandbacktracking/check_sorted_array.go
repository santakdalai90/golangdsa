package recursionandbacktracking

// IsSorted checks if an input array is sorted using recursion
func IsSorted(arr []int) bool {
	if len(arr) <= 1 {
		return true
	}

	return arr[0] < arr[1] && IsSorted(arr[1:])
}
