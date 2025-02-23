package logic_02

func SoalNo10(n int) (result [][]int) {
	result = make([][]int, n)

	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
		start := 1

		for j := 0; j <= i; j++ {
			result[i][j] = start
			start += 2
		}
	}
	return result
}
