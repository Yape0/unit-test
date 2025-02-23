package logic_02

func SoalNo5(n int) (result [][]int) {
	start := 1
	result = make([][]int, n)

	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
		if i%2 == 0 {
			for j := 0; j < n; j++ {
				result[i][j] = start
				start += 2
			}
		} else {
			for j := 0; j < n; j++ {
				result[i][n-1-j] = start
				start += 2
			}
		}

	}
	return result
}
