package main

import (
	"fmt"
	"testing"
)

func Test_isAcronym(t *testing.T) {
	fmt.Println("")
}

// leetcode submit region begin(Prohibit modification and deletion)
func isAcronym(words []string, s string) bool {
	if len(words) != len(s) {
		return false
	}
	for i := 0; i < len(words); i++ {
		if s[i] != words[i][0] {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
