package piscine

import "github.com/01-edu/z01"

func PrintStr(str string) {
	stringnew := []rune(str)

	for i := range stringnew {
		z01.PrintRune(stringnew[i])
	}
}

func main() {
	str := "Hello World!"
	PrintStr(str)
}
