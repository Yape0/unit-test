package logic_03

func SoalNo04(n int) [][]int {
	result := make([][]int, n)
	start := 1
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
		geser := 0
		for j := 0; j < n; j++ {
			if i+j >= n-1 {
				if i%2 == 1 {
					result[i][j] = start
				} else {
					//seharusnya 0.8
					//seharusnya yang diisi 6,7,8
					result[i][n-1-geser] = start
					//result[i][j] = start
					geser++

				}
				start += 2
			}
		}
	}
	return result
}
