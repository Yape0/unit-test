package logic_01

func Soal_06(n int) []int {
	slice := make([]int, n)
	start := 30

	for i := 0; i < n; i++ {
		slice[i] = start
		start -= 3
	}
	return slice
}
