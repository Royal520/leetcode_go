package cn

import (
	"fmt"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 1, 2, 3, 3}))
}

// leetcode submit region begin(Prohibit modification and deletion)
func removeDuplicates(nums []int) int {
	ans := 0
	count := 1
	for _, value := range nums[1:] {
		if value != nums[ans] {
			ans++
			nums[ans] = value
			count = 1
		} else {
			if count == 1 {
				ans++
				nums[ans] = value
				count = 2
			}
		}
	}
	return ans + 1
}

//leetcode submit region end(Prohibit modification and deletion)
