package cn

import (
	"fmt"
	"testing"
)

func Test_SortAnArray(t *testing.T) {
	a := sortArray([]int{5, 2, 3, 1, 8, 0})
	fmt.Print(a)
}

// leetcode submit region begin(Prohibit modification and deletion)
func sortArray(nums []int) []int {
	var quick func(nums []int, left, right int) []int
	quick = func(nums []int, left, right int) []int {
		// 递归终止条件
		if left > right {
			return nil
		}
		// 左右指针及主元
		i, j, pivot := left, right, nums[left]
		for i < j {
			// 寻找小于主元的右边元素
			for i < j && nums[j] >= pivot {
				j--
			}
			// 寻找大于主元的左边元素
			for i < j && nums[i] <= pivot {
				i++
			}
			// 交换i/j下标元素
			nums[i], nums[j] = nums[j], nums[i]
		}
		// 交换元素
		nums[i], nums[left] = nums[left], nums[i]
		quick(nums, left, i-1)
		quick(nums, i+1, right)
		return nums
	}
	return quick(nums, 0, len(nums)-1)
}

//leetcode submit region end(Prohibit modification and deletion)
