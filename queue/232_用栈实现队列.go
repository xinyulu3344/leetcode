package queue

// https://leetcode.cn/problems/implement-queue-using-stacks/

type MyQueue struct {
	in []int
	out []int
}


func Constructor() MyQueue {
	return MyQueue{
		in: make([]int, 0),
		out: make([]int, 0),
	}
}


func (this *MyQueue) Push(x int)  {
	this.in = append(this.in, x)
}


func (this *MyQueue) Pop() int {
	if this.Empty() {
		return -1
	}
	if len(this.out) != 0 {
		r := this.out[len(this.out) - 1]
		this.out = this.out[:len(this.out) - 1]
		return r
	}
	for i := len(this.in) - 1; i >= 0; i-- {
		this.out = append(this.out, this.in[i])
	}
	this.in = this.in[:0]
	r := this.out[len(this.out) - 1]
	this.out = this.out[:len(this.out) - 1]
	return r
}


func (this *MyQueue) Peek() int {
	if this.Empty() {
		return -1
	}
	if len(this.out) != 0 {
		return this.out[len(this.out) - 1]
	}
	for i := len(this.in) - 1; i >= 0; i-- {
		this.out = append(this.out, this.in[i])
	}
	this.in = this.in[:0]
	return this.out[len(this.out) - 1]
}


func (this *MyQueue) Empty() bool {
	if len(this.in) == 0 && len(this.out) == 0 {
		return true
	}
	return false
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */