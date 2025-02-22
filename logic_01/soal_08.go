package logic_01

func Soal_08(n int) []int {
	slice := make([]int, n)

	start := 2
	median := 0.5 * float32(n)
	for i := 1; i <= n; i++ {
		slice[i-1] = start
		if float32(i) < median {
			start += 2
		} else if float32(i) > median {
			start -= 2
		}
	}
	return slice
}
