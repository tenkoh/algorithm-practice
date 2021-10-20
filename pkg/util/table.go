package util

func MakeTable(row, col, init int) [][]int {
	ret := make([][]int, row)
	for r := range ret {
		ret[r] = make([]int, col)
	}
	if init == 0 {
		return ret
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			ret[i][j] = init
		}
	}
	return ret
}
