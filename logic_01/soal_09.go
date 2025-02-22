package logic_01

func Soal_09(n int) []int {
	slice := make([]int, n)

	start := 3
	median := 0.5 * float32(n)
	for i := 1; i <= n; i++ {
		slice[i-1] = start
		if float32(i) < median {
			start += 3
		} else if float32(i) > median {
			start -= 3
		}
	}
	return slice
}
