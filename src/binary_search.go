package src

import (
    log "github.com/sirupsen/logrus"
    "os"
)

func init(){
     // Log as JSON instead of the default ASCII formatter.
  log.SetFormatter(&log.JSONFormatter{})

  // Output to stdout instead of the default stderr
  // Can be any io.Writer, see below for File example
  log.SetOutput(os.Stdout)

  // Only log the warning severity or above.
  log.SetLevel(log.DebugLevel)
}

// BinarySearchIterative decides if the binary search has to happen recursive or iterative
var BinarySearchIterative=true

// BinarySearch searches for an integer given an integer array. 
// It returns the position of the element if found, else -1
func BinarySearch(arr []int, elem int) int {
    if BinarySearchIterative {
        return binarySearchIterative(arr, elem)
    }
    return binarySearchRecursive(arr, 0, len(arr)-1, elem)
}

func binarySearchRecursive(arr []int, low, high, elem int) int {
    if low > high {
        return -1
    }
    mid := (low + high)/2
    if arr[mid] == elem {
        return mid
    }
    if arr[mid] > elem {
        return binarySearchRecursive(arr, low, mid-1, elem)
    } else {
        return binarySearchRecursive(arr, mid+1, high, elem)
    }
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