package logic_02

func SoalNo9(n int) (result [][]int) {
	result = make([][]int, n)
	start := 1
	count := 8

	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		result[i][i] = start
		result[count][i] = start
		start += 2
		count--
	}
	return result
}
