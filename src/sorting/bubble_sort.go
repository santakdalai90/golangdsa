package sorting

// BubbleSort sorts and array using bubble sort
// bubble sort makes sure that in ith iteration
// the ith biggest number is at its position
func BubbleSort(arr []int) {
	// to track the last index to be evaluated
	// as in each iteration we need to check one less item
	lastIdx := len(arr) - 1
	// loop only till all are not sorted
	sorted := false
	for !sorted {
		sorted = true
		for i := 0; i < lastIdx; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
                sorted = false
			}
		}
		lastIdx--
	}
}
