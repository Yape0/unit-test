package logic_02

func SoalNo11(n int) (result [][]int) {
	result = make([][]int, n)

	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
		start := 17
		for j := 8; j >= i; j-- {
			result[i][j] = start
			start -= 2
		}
	}
	return result
}
