package recursion_and_backtracking

// Generate all the strings of length n from a given set of k characters
func KAryString(chars []byte, n int) []string{
    if n == 0 {
        return []string{""}
    }
    ans := make([]string, 0)
    prevStrings := KAryString(chars, n-1)
    for _, s := range prevStrings{
        for _, c := range chars {
            ans = append(ans, string(c)+s)
        }
    }
    return ans
}