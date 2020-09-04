package _55_minstack

type MinStack struct {
	stack  []int
	minVal int
	isInit bool
}


func Constructor() MinStack {
	return MinStack{
		stack: make([]int, 0),
	}
}


func (this *MinStack) Push(x int)  {
	if !this.isInit {
		this.minVal = x
		this.stack = append(this.stack, 0)
		this.isInit = true
		return
	}
	delta := x - this.minVal
	if delta < 0 {
		this.minVal = x
	}
	this.stack = append(this.stack, delta)
}


func (this *MinStack) Pop()  {
	val := this.stack[len(this.stack) - 1]
	if val < 0 {
		this.minVal = this.minVal - val
	}
	this.stack = this.stack[:len(this.stack) - 1]
}


func (this *MinStack) Top() int {
	val := this.stack[len(this.stack) - 1]
	return val + this.minVal
}


func (this *MinStack) GetMin() int {
	return this.minVal
}