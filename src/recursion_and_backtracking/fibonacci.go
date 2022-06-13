package recursion_and_backtracking

var mem = make(map[int]int)    //memoization
// Fibonacci returns nth Fibonacci number of a series which starts like
// 0, 1, 1, 2, 3, 5, 8, ...
func Fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    if v, ok := mem[n]; ok {
        return v
    }
    mem[n] = Fibonacci(n-1) + Fibonacci(n-2)
    return mem[n]
}

var memIter = make(map[int]int)    //memoization
func FibonacciIterative(n int) int {
    if n <= 1 {
        return n
    }
    if v, ok := memIter[n]; ok {
        return v
    }
    a,b := 0, 1
    var val int
    for i := 2; i <= n; i++ {
        val = a + b
        a = b
        b = val
    }
    mem[n] = val
    return mem[n]
}