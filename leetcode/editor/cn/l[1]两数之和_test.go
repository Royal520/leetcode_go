package cn

import (
	"fmt"
	"testing"
)

func Test_TwoSum(t *testing.T) {
	fmt.Println("lovejj")
}

// leetcode submit region begin(Prohibit modification and deletion)
func twoSum(nums []int, target int) []int {
	idx := map[int]int{}     // 创建一个空哈希表
	for j, x := range nums { // 枚举 j
		// 在左边找 nums[i]，满足 nums[i]+x=target
		if i, ok := idx[target-x]; ok { // 找到了
			return []int{i, j} // 返回两个数的下标
		}
		idx[x] = j // 保存 nums[j] 和 j
	}
	return nil
}

//leetcode submit region end(Prohibit modification and deletion)
