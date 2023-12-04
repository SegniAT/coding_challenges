package main

type Item struct {
	val int
	min int
}

type MinStack struct {
	stack []Item
}

func Constructor() MinStack {
	return MinStack{
		stack: []Item{},
	}
}

func (this *MinStack) Push(val int) {
	if this.IsEmpty() {
		this.stack = append(this.stack, Item{val: val, min: val})
	} else {
		prevMin := this.stack[len(this.stack)-1].min
		if val < prevMin {
			this.stack = append(this.stack, Item{val: val, min: val})
		} else {
			this.stack = append(this.stack, Item{val: val, min: prevMin})
		}
	}
}

func (this *MinStack) Pop() {
	if this.IsEmpty() {
		return
	}

	this.stack = this.stack[0 : len(this.stack)-1]
}

func (this *MinStack) Top() int {
	if this.IsEmpty() {
		return -1
	}

	return this.stack[len(this.stack)-1].val
}

func (this *MinStack) GetMin() int {
	if this.IsEmpty() {
		return -1
	}

	return this.stack[len(this.stack)-1].min
}

func (this *MinStack) IsEmpty() bool {
	return len(this.stack) == 0
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main() {

}
