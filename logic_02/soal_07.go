package logic_02

func SoalNo7(n int) (result [][]int) {
	result = make([][]int, n)
	start := 1

	for j := 0; j < n; j++ {
		result[j] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		result[i][i] = start
		start += 2
	}
	return result
}
