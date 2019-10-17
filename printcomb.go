package piscine

import "github.com/01-edu/z01"

func PrintComb() {

	var a, b, c rune = '0', '1', '2'
	for i := a; i <= '7'; i++ {
		for j := b; j <= '8'; j++ {
			if j > i {
				for k := c; k <= '9'; k++ {
					if k > j {
						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(k)
						if i == '7' && j == '8' && k == '9' {
							z01.PrintRune('\n')
						} else {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	}
}
