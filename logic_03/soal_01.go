package logic_03

func SoalNo01(n int) (result [][]int) {
	result = make([][]int, n)
	start := 1

	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
		if i%2 == 0 {
			// For even rows, fill left to right with increasing odd numbers
			for j := 0; j <= i; j++ {
				result[i][j] = start
				start += 2
			}
		} else {
			// For odd rows, fill left to right with decreasing odd numbers
			number := start + 2*i
			for j := 0; j <= i; j++ {
				result[i][j] = number
				number -= 2
			}
			start += 2 * (i + 1)
		}
	}

	return result
}
