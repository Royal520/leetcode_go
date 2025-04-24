package main

import (
	"fmt"
	"testing"
)

func Test_HouseRobber(t *testing.T) {
	fmt.Println(rob([]int{226, 174, 214, 16, 218, 48, 153, 131, 128, 17, 157, 142, 88, 43, 37, 157, 43, 221, 191, 68, 206, 23, 225, 82, 54, 118, 111, 46, 80, 49, 245, 63, 25, 194, 72, 80, 143, 55, 209, 18, 55, 122, 65, 66, 177, 101, 63, 201, 172, 130, 103, 225, 142, 46, 86, 185, 62, 138, 212, 192, 125, 77, 223, 188, 99, 228, 90, 25, 193, 211, 84, 239, 119, 234, 85, 83, 123, 120, 131, 203, 219, 10, 82, 35, 120, 180, 249, 106, 37, 169, 225, 54, 103, 55, 166, 124}))
}

// leetcode submit region begin(Prohibit modification and deletion)
func maxN(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func rob(nums []int) int {
	prev := 0
	curr := 0
	// 每次循环，计算“偷到当前房子为止的最大金额”
	for _, i := range nums {
		// 循环开始时，curr 表示 dp[k-1]，prev 表示 dp[k-2]
		// dp[k] = max{ dp[k-1], dp[k-2] + i }
		temp := maxN(curr, prev+i)
		prev = curr
		curr = temp
		// 循环结束时，curr 表示 dp[k]，prev 表示 dp[k-1]
	}

	return curr
}

//leetcode submit region end(Prohibit modification and deletion)
