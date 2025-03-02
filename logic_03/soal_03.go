package logic_03

func SoalNo03(n int) [][]int {
	result := make([][]int, n)
	start := 1

	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i <= j {
				if i == 0 {
					result[i][j-i] = start
					start += 3
				} else if i%2 == 0 {
					result[i][j-i] = start
					start += 2
					if j == n-1 {
						start += 1
					}
				} else {
					result[i][n-j-1] = start
					start += 3
					if j == n-1 {
						start += 1
					}
				}
			}

		}
	}

	return result
}
