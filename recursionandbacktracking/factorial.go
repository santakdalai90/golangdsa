package recursionandbacktracking

func Factorial(n int) int {
	if n <= 1 {
		return n
	}
	return n * Factorial(n-1)
}

func FactorialIterative(n int) int {
	ans := 1
	for ; n >= 1; n-- {
		ans *= n
	}
	return ans
}
