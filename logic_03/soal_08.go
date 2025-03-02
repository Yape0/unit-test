package logic_03

func SoalNo08(n int) [][]int {
	result := make([][]int, n)
	mid := (n - 1) / 2
	start := 1
	for i := range result {
		result[i] = make([]int, n)
	}

	for k := 0; k <= mid; k++ {
		for b := 0; b < n; b++ {
			if k+b >= mid && b-k <= mid {
				if k%2 == 1 {
					result[b][k] = start
					result[b][n-1-k] = start
				} else {
					result[n-1-b][k] = start
					result[n-1-b][n-1-k] = start
				}
				start += 2
			}

		}

	}
	return result
}
