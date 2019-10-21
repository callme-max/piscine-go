package piscine

func IterativeFactorial(nb int) int {
	result := 1

	if nb < 0 {
		result = 0
		return result
	}

	if nb >= 0 && nb < 18 {
		for i := 1; i < nb; i++ {
			result = result * (i + 1)
		}
		return result
	}
	result = 0
	return result
}
