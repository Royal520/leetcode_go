package cn

import (
	"fmt"
	"testing"
)

func Test_MinStack(t *testing.T) {
	fmt.Println("lovejj")
}

// leetcode submit region begin(Prohibit modification and deletion)
type MinStack struct {
	Min, Arr []int
}

//func Constructor() MinStack {
//	return MinStack{}
//}

func (p *MinStack) Push(val int) {
	p.Arr = append(p.Arr, val)
	if len(p.Min) == 0 || val < p.Min[len(p.Min)-1] {
		p.Min = append(p.Min, val)
		return
	}
	p.Min = append(p.Min, p.Min[len(p.Min)-1])
}

func (p *MinStack) Pop() {
	p.Arr = p.Arr[:len(p.Arr)-1]
	p.Min = p.Min[:len(p.Min)-1]
}

func (p *MinStack) Top() int {
	return p.Arr[len(p.Arr)-1]
}

func (p *MinStack) GetMin() int {
	return p.Min[len(p.Min)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
//leetcode submit region end(Prohibit modification and deletion)
