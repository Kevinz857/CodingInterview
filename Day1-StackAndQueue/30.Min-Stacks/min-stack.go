type MinStack struct {
	S []int
	// 单调递减栈
	MinS []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	// 初始化
	return MinStack{
		S:    make([]int, 0, 16),
		MinS: make([]int, 0, 16),
	}
}

func (this *MinStack) Push(x int) {
	this.S = append(this.S, x)

	// 如果单调递减栈中无元素，或者栈顶元素大于等于插入元素，就入到单调栈中
	if len(this.MinS) == 0 || x <= this.MinS[len(this.MinS)-1] {
		this.MinS = append(this.MinS, x)
	}
}

func (this *MinStack) Pop() {
	if len(this.S) < 1 {
		return
	}
	tm := this.S[len(this.S)-1]
	this.S = this.S[:len(this.S)-1]
	// 如果出栈元素和单调栈元素相等，也出栈
	if tm == this.MinS[len(this.MinS)-1] {
		this.MinS = this.MinS[:len(this.MinS)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.S) < 1 {
		return -1
	}
	return this.S[len(this.S)-1]
}

func (this *MinStack) Min() int {
	if len(this.MinS) == 0 {
		return -1
	}
	// 返回单调栈栈顶元素
	return this.MinS[len(this.MinS)-1]
}

/*
package leetcode

// MinStack define
type MinStack struct {
	elements, min []int
	l             int
}

// initialize your data structure here.

// Constructor155 define
func Constructor155() MinStack {
	return MinStack{make([]int, 0), make([]int, 0), 0}
}

// Push define
func (this *MinStack) Push(x int) {
	this.elements = append(this.elements, x)
	if this.l == 0 {
		this.min = append(this.min, x)
	} else {
		min := this.GetMin()
		if x < min {
			this.min = append(this.min, x)
		} else {
			this.min = append(this.min, min)
		}
	}
	this.l++
}

func (this *MinStack) Pop() {
	this.l--
	this.min = this.min[:this.l]
	this.elements = this.elements[:this.l]
}

func (this *MinStack) Top() int {
	return this.elements[this.l-1]
}

func (this *MinStack) GetMin() int {
	return this.min[this.l-1]
}
*/

/**
* Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
*/
