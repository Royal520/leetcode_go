package main

import (
	"math"
	"strconv"
	"testing"
)

func Test_MaximumValueOfAStringInAnArray(t *testing.T) {
	print(maximumValue([]string{"alic3", "bob", "3", "4", "00000"}))
}

// leetcode submit region begin(Prohibit modification and deletion)
func maximumValue(strs []string) int {
	maxLen := 0
	for _, str := range strs {
		atoi, err := strconv.Atoi(str)
		if err != nil {
			maxLen = int(math.Max(float64(maxLen), float64(len(str))))
		}
		maxLen = int(math.Max(float64(maxLen), float64(atoi)))
	}
	return maxLen
}

//leetcode submit region end(Prohibit modification and deletion)
