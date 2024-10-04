package recursionandbacktracking

func ReverseString(s string) string {
	if len(s) <= 1 {
		return s
	}
	return ReverseString(s[1:]) + string(s[0])
}
