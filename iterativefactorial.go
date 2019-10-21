package piscine

func IterativeFactorial(nb int) int {
	result := 1

	if nb < 0 {
		result = 0
		return result
	}
	for i := 1; i < nb; i++ {
		result = result * (i + 1)
	}
	if result > 0 && nb < 2000000 {
		return result
	}
	result = 0
	return result
}
