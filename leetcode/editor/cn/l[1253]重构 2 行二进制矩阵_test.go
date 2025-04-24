package main

import (
	"fmt"
	"testing"
)

func Test_ReconstructA2RowBinaryMatrix(t *testing.T) {
	matrix := reconstructMatrix(9, 2, []int{0, 1, 2, 0, 0, 0, 0, 0, 2, 1, 2, 1, 2})
	fmt.Println(matrix)
}

// leetcode submit region begin(Prohibit modification and deletion)
func reconstructMatrix(upper int, lower int, colsum []int) [][]int {
	res := make([][]int, 2)
	sum := 0
	a, b := upper, lower
	// 第2步，内层数组初始化
	for i := 0; i < 2; i++ {
		res[i] = make([]int, len(colsum))
	}
	for i, v := range colsum {
		if v == 1 {
			if a > b {
				res[0][i] = 1
				a--
			} else {
				res[1][i] = 1
				b--
			}
		}
		if v == 2 {
			if lower == 0 || upper == 0 {
				return res[0:0]
			}
			res[0][i] = 1
			a--
			res[1][i] = 1
			b--
		}
		sum += colsum[i]
	}
	if sum != lower+upper || a < 0 || b < 0 {
		return res[0:0]
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
