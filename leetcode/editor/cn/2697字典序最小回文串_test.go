package main

import (
	"testing"
)

func Test_makeSmallestPalindrome(t *testing.T) {
	println("hello world                                 ")
}

// leetcode submit region begin(Prohibit modification and deletion)
func makeSmallestPalindrome(s string) string {
	left, right := 0, len(s)-1
	t := []byte(s)
	var t2 int = 9
	print(t2)
	for left < right {
		if s[left] != s[right] {
			t[left] = byte(min(int(s[left]), int(s[right])))
			t[right] = t[left]
		}
		left++
		right--
	}
	return string(t)
}

//leetcode submit region end(Prohibit modification and deletion)
