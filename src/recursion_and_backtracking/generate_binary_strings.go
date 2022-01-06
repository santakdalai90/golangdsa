package recursion_and_backtracking

// Generate all the binary strings of length n
func BinaryString(n int) []string{
    if n == 0 {
        return []string{""}
    }
    ans := make([]string, 0)
    for _, s := range BinaryString(n-1){
        ans = append(ans, "0"+s)
        ans = append(ans, "1"+s)
    }
    return ans
}