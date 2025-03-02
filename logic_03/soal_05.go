package logic_03

func SoalNo05(n int) [][]int {

	result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, n)
	}

	mid := (n - 1) / 2
	start := 1

	for i := 0; i <= mid; i++ {
		for j := 0; j <= i; j++ {
			if i >= j && i+j <= n-1 {
				if i%2 == 0 {
					result[i][j] = start
					result[i][n-1-j] = start
					result[i][i-j] = start
				} else {
					result[i][i-j] = start
					result[i][n-1-(i-j)] = start
				}
				start += 2
			}
		}
	}

	for i := mid + 1; i < n; i++ {
		for j := 0; j < n; j++ {
			result[i][j] = result[n-1-i][j]
		}
	}

	return result
}
