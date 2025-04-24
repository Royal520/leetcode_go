package cn

import (
	"testing"
)

func Test_NumberOfIslands(t *testing.T) {
	arry := [][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}}
	println(numIslands(arry))
}

// leetcode submit region begin(Prohibit modification and deletion)
type index struct {
	i, j int
}

func numIslands(grid [][]byte) int {
	return 1
}

//leetcode submit region end(Prohibit modification and deletion)
