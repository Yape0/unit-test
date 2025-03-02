package logic_03

func SoalNo07(n int) [][]int {
	result := make([][]int, n)
	mid := (n - 1) / 2
	start := 1
	for i := range result {
		result[i] = make([]int, n)
	}

	for row := 0; row <= mid; row++ {
		for col := 0; col < n; col++ {
			if row+col >= mid && col-row <= mid {
				if row%2 == 1 {
					result[row][col] = start
					result[n-1-row][col] = start
				} else {
					result[row][n-1-col] = start
					result[n-1-row][n-1-col] = start
				}
				start += 2
			}
		}
	}
	return result
}
