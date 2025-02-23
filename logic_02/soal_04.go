package logic_02

func SoalNo4(n int) (result [][]int) {
	result = make([][]int, n)

	for i := range result {
		result[i] = make([]int, n)
	}

	start := 1
	increment := 3
	count := 1

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			result[i][j] = start
			start += increment
			count++
			if count == 11 {
				count = 0
				if increment == 3 {
					increment = 2
				} else {
					increment = 3
				}
			}
		}
	}
	return result
}
