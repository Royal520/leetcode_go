package main

import (
	"fmt"
	"math"
	"testing"
)

func Test_FindThePivotInteger(t *testing.T) {
	fmt.Println("lovejj")
}

// leetcode submit region begin(Prohibit modification and deletion)
func pivotInteger(n int) int {
	t := (n*n + n) / 2
	x := int(math.Sqrt(float64(t)))
	if x*x == t {
		return x
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
