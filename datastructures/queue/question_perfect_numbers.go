package queue

/*
A perfect number is defined as a number consisting of only 1 and 2 and is an even length palindrome.
e.g.
11, 22, 1111, 1221, 2112, 2222, ....
Find nth perfect number
*/

func reverseString(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func GetPerfectNumber(n int) string {
	q := New()

	q.Push("1")
	q.Push("2")

	var f interface{}
	for i := 0; i < n; i++ {
		f, _ = q.Pop()
		q.Push(f.(string) + "1")
		q.Push(f.(string) + "2")
	}

	return f.(string) + reverseString(f.(string))
}
