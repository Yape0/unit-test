package logic_02

func SoalNo1(n int) (result [][]int) {
	result = make([][]int, n)

	for i := 0; i < n; i++ {

		result[i] = make([]int, n)
		num := 1

		for j := 0; j < n; j++ {
			result[i][j] = num
			num += 2
		}

	}
	return result
}
