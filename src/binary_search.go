package src

// BinarySearchIterative decides if the binary search has to happen recursive or iterative
var BinarySearchIterative=true

// BinarySearch searches for an integer given an integer array. 
// It returns the position of the element if found, else -1
func BinarySearch(arr []int, elem int) int {
    if BinarySearchIterative {
        return binarySearchIterative(arr, elem)
    }
    return binarySearchRecursive(arr, elem)
}

func binarySearchRecursive(arr []int, elem int) int {
    return -1
}

func binarySearchIterative(arr []int, elem int) int {
    low := 0
    high := len(arr) - 1

    for low <= high {
        mid := (low + high)/2
        if arr[mid] == elem {
            return mid
        }

        if arr[mid] > elem {
            high = mid - 1
        } else {
            low = mid + 1
        }
    }

    return -1
}