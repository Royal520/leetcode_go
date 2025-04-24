package cn

import (
	"math"
	"sort"
	"testing"
)

func Test_LongestConsecutiveSequence(t *testing.T) {
	print(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
}

// leetcode submit region begin(Prohibit modification and deletion)
func longestConsecutive(nums []int) int {
	sort.Ints(nums)
	if len(nums) < 2 {
		return len(nums)
	}
	maxLen := 1
	start := 0
	tempLen := 1
	for i := 1; i < len(nums); i++ {
		dis := nums[i] - nums[start]
		if dis == 1 {
			tempLen++
		} else if dis > 1 {
			tempLen = 1
		}
		start = i
		maxLen = int(math.Max(float64(tempLen), float64(maxLen)))
	}
	return maxLen
}

//leetcode submit region end(Prohibit modification and deletion)
