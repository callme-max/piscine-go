package piscine

func StrLen(str string) int {
	str_length := []rune(str)
	nb := 0
	for range str_length {
		nb++
	}
	return nb
}
