package sorting

func InsertionSort(arr []int) {
    if len(arr) <= 1 {
        return
    }
    InsertionSort(arr[:len(arr)-1])
    pos := len(arr) - 1
    for i := len(arr)-2; i>=0; i-- {
        if arr[i] > arr[pos] {
            arr[i], arr[pos] = arr[pos], arr[i]
            pos --
        } else {
            break
        }
    }
}