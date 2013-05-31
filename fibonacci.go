package fibonacci

func Calclulate(n int) (result int) {
	if n < 2 {
		result = 1
	} else {
		result = Calclulate(n-2) + Calclulate(n-1)
	}
	return result
}
