package logic_02

func SoalNo3(n int) (result [][]int) {
	result = make([][]int, n)
	num := 1
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)

		for j := 0; j < n; j++ {
			result[i][j] = num
			num += 2
		}

	}
	return result
}
