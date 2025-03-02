package logic_02

func SoalNo12kaku(n int) (result [][]int) {
	result = make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
		start := 1
		for j := 0; j < n; j++ {
			if j >= 1 && j <= n-1 {
				result[i][j] = start
				start += 2
			} else if j >= n-i-1 && j <= i {
				result[i][j] = start
				start += 2
			} else {
				result[i][j] = 0
			}
		}
	}
	return result
}

func SoalNo12(n int) (result [][]int) {
	result = make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
		start := 1
		for j := 0; j < n; j++ {
			if j >= 1 && j <= n-1 {
				result[i][j] = start
				start += 2
			} else if j >= n-i-1 && j <= i {
			}
		}
	}
}
