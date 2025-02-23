package logic_02

func SoalNo8(n int) (result [][]int) {
	result = make([][]int, n)
	start := 17
	count := 8

	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}
	for j := 0; j < n; j++ {
		result[j][count] = start
		start -= 2
		count--

	}
	return result
}
