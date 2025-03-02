package logic_03

func SoalNo02(n int) [][]int {
	result := make([][]int, n)

	start := 1

	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
		for j := i; j < n; j++ {
			if i%2 == 0 {
				result[i][j] = start
				start += 2
			} else {
				result[i][n-1-(j-i)] = start
				start += 2
			}
		}
	}
	return result
}
