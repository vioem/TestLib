package pkg

func SumSeries(a, b int) int {
	var result int
	for i := a; i <= b; i++ {
		result += i
	}
	return result
}
