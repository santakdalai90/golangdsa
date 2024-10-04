package sorting

func merge(a, b []int) []int {
    arr := make([]int, 0)
    p,q := 0, 0
    for p<len(a) && q<len(b) {
        if a[p] < b[q] {
            arr = append(arr, a[p])
            p++
        } else {
            arr = append(arr, b[q])
            q++
        }
    }
    
    if p < len(a) {
        arr = append(arr, a[p:]...)
    } else if q < len(b) {
        arr = append(arr, b[q:]...)
    }

    return arr
}

func MergeSort(arr []int) {
    l := len(arr)
    if l == 1 {
        return
    }
    a := arr[:l/2]
    b := arr[l/2:]
    MergeSort(a)
    MergeSort(b)
    arr = merge(a, b)
}