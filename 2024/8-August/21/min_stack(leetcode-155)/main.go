package main

func main() {
}

// val: actual value
// min: minimum value so far
type Value struct {
	val int
	min int
}

type MinStack struct {
	stack []Value
}

func Constructor() MinStack {
	return MinStack{stack: []Value{}}
}

func (this *MinStack) Push(val int) {
	if N := len(this.stack); N > 0 {
		prevMin := this.stack[N-1].min
		this.stack = append(this.stack, Value{val: val, min: min(prevMin, val)})
	} else {
		this.stack = append(this.stack, Value{val: val, min: val})
	}
}

func (this *MinStack) Pop() {
	N := len(this.stack)
	this.stack = this.stack[:N-1]
}

func (this *MinStack) Top() int {
	N := len(this.stack)
	return this.stack[N-1].val
}

func (this *MinStack) GetMin() int {
	N := len(this.stack)
	return this.stack[N-1].min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
