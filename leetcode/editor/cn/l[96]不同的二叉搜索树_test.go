package main

import (
	"fmt"
	"testing"
)

func Test_UniqueBinarySearchTrees(t *testing.T) {
	fmt.Println(numTrees(4))
}

// leetcode submit region begin(Prohibit modification and deletion)
func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] += dp[i-1] * 2
		for j := 2; j <= i/2; j++ {
			dp[i] += dp[j-1] * dp[i-j] * 2
		}
		if i%2 == 1 {
			dp[i] += dp[i/2] * dp[i-i/2-1]
		}
	}
	return dp[n]
}

//leetcode submit region end(Prohibit modification and deletion)
