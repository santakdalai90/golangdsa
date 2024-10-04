package util

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func Min(args ...int) int {
	if len(args) == 0 {
		panic("no arguments for Min funtion")
	}
	min := args[0]
	for _, x := range args {
		if x < min {
			min = x
		}
	}
	return min
}

func IntCompare(a, b int) int {
	return a - b
}

func Max[T any](compare func(a, b T) int, elements ...T) T {
	if len(elements) == 0 {
		panic("no max for empty arguments")
	}
	var ans T = elements[0]
	for _, x := range elements {
		if compare(ans, x) > 0 { //ans > x
			ans = x
		}
	}
	return ans
}
