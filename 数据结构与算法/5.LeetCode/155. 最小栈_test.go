package leetcode

import "container/list"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

// min栈谁小压谁， 同步弹出
type MinStack struct {
	data *list.List // 数据栈
	min  *list.List // 最小值栈
}

func Constructor() MinStack {
	return MinStack{
		data: list.New(),
		min:  list.New(),
	}
}

func (this *MinStack) Push(val int) {
	this.data.PushBack(val)
	if this.min.Len() == 0 {
		this.min.PushBack(val)
	} else {
		top := this.min.Back().Value.(int)
		if top < val {
			this.min.PushBack(top)
		} else {
			this.min.PushBack(val)
		}
	}
}

func (this *MinStack) Pop() {
	this.data.Remove(this.data.Back())
	this.min.Remove(this.min.Back())
}

func (this *MinStack) Top() int {
	return this.data.Back().Value.(int)
}

func (this *MinStack) GetMin() int {
	return this.min.Back().Value.(int)
}
