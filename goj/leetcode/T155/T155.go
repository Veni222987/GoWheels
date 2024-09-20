package main

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:    make([]int, 0),
		minStack: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if len(this.minStack) == 0 {
		this.minStack = []int{val}
	} else {
		var i int
		for i = 0; i < len(this.minStack); i++ {
			if this.minStack[i] >= val {
				break
			}
		}
		this.minStack = append(this.minStack[0:i], append([]int{val}, this.minStack[i:]...)...)
	}

}

func (this *MinStack) Pop() {
	popVal := this.stack[len(this.stack)-1]
	this.stack = this.stack[0 : len(this.stack)-1]
	for i := range this.minStack {
		if this.minStack[i] == popVal {
			if i == len(this.minStack)-1 {
				this.minStack = this.minStack[0:i]
			} else {
				this.minStack = append(this.minStack[0:i], this.minStack[i+1:]...)
			}
			return
		}
	}
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[0]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
