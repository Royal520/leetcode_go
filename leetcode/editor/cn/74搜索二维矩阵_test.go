package cn

import (
	"fmt"
	"sort"
	"testing"
)

func Test_searchMatrix(t *testing.T) {
	fmt.Println("")
}

// leetcode submit region begin(Prohibit modification and deletion)
func searchMatrix(matrix [][]int, target int) bool {
	row := sort.Search(len(matrix), func(i int) bool { return matrix[i][0] > target }) - 1
	if row < 0 {
		return false
	}
	col := sort.SearchInts(matrix[row], target)
	return col < len(matrix[row]) && matrix[row][col] == target
}

//leetcode submit region end(Prohibit modification and deletion)
