package pow

func Power(i int) int {
	if i < 2 {
		return 1
	}
	return i * Power(i-1)
}
