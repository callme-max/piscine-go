package piscine

func StrRev(s string) string {
	var n int = StrLen(s)
	string_new := []rune(s)
	s = ""

	for i := 0; i < n; i++ {
		s = s + string(string_new[n-i-1])
	}
	return s
}
