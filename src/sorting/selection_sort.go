package sorting

// SelectionSort sorts a given array using selection sort.
// Each iteration finds the minimum and swaps it to its position
// Loop runs till the second last position. Last position would be automatically
// done in the last iteration
func SelectionSort(arr []int) {
    for i:=0; i<len(arr)-1; i++ {
        minIdx := i
        // find index of the minimum
        for j:=i+1; j< len(arr); j++ {
            if arr[j] < arr[minIdx] {
                minIdx = j
            }
        }
        if minIdx != i {        // avoid swap if current number is minimum
            // swap the minimum to its position
            arr[i], arr[minIdx] = arr[minIdx], arr[i]
        }
    }
}