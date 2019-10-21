package piscine

func IterativeFactorial(nb int) int {
	result := 1

	for i := 1; i < nb; i++ {
		result = result * (i + 1)
	}
	return result
}
