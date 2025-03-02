package logic_03

func SoalNo14(n int) [][]int {
	result := make([][]int, n)
	start := 1
	for i := range result {
		result[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		num := start
		for j := 0; j < n; j++ {
			result[i][j] = num
			num += 34
		}
		start += 2
	}
	return result
}
