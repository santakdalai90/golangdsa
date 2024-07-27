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