package logic_03

func SoalNo13(n int) [][]int {
	result := make([][]int, n)

	mid := (n - 1) / 2
	for i := range result {
		result[i] = make([]int, n)
	}

	for i := 0; i <= mid; i++ {
		start := 1
		for j := 0; j <= mid; j++ {
			if i+j >= mid {
				if i%2 == 0 && j%2 == 0 || i%2 != 0 && j%2 != 0 {
					result[i][j] = start
					result[n-1-i][j] = start
					result[n-1-i][n-1-j] = start
					result[i][n-1-j] = start
				}
			}
			start += 2
		}
	}
	return result
}
