package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

type Vector2D struct {
	// 构造函数将把所有数字放入这个列表中。
	vector [][]int
	inner  int
	outer  int
}

func Constructor(vec [][]int) Vector2D {
	vector := vec
	return Vector2D{vector: vector}
}

func (this *Vector2D) advanceToNext() {
	// 虽然 outer 仍然在向量内，但 internal 位于 outer 指向的 internal 列表的末尾，
	// 我们希望向前移动到下一个 internal 向量的起点。
	for this.outer < len(this.vector) && this.inner == len(this.vector[this.outer]) {
		this.inner = 0
		this.outer++
	}
}

func (this *Vector2D) Next() int {
	this.advanceToNext()
	// 返回当前元素并向内移动，使其位于当前元素之后。
	result := this.vector[this.outer][this.inner]
	this.inner++
	return result
}

func (this *Vector2D) HasNext() bool {
	// 确保移动位置指针以使其指向整数，
	// 或者让 outer = vector.length.
	this.advanceToNext()
	// 如果 outer = vector.length 那么就没有剩下的整数了,
	// 否则我们会在一个整数停下，也就是还有剩下的整数。
	return this.outer < len(this.vector)
}
