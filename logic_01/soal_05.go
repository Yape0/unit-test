package logic_01

func Soal_05(n int) []int {

	slice := make([]int, n)
	num := 20
	for i := 0; i < n; i++ {
		slice[i] = num
		num -= 2
	}
	return slice
}
