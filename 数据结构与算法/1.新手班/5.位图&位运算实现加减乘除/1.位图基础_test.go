package class05

import "fmt"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	bitMap := NewBitMap(1024)
	bitMap.Add(1024)
	fmt.Println(bitMap.Find(1024))
	bitMap.Add(1023)
	fmt.Println(bitMap.Find(1023))
	bitMap.Add(0)
	fmt.Println(bitMap.Find(0))
	bitMap.Add(1)
	fmt.Println(bitMap.Find(1))
	bitMap.Delete(1)
	fmt.Println(bitMap.Find(1))

	return "Hello World!", nil
}

func NewBitMap(maxVal int) *BitMap {
	return &BitMap{
		data:   make([]int64, (maxVal/64)+1),
		maxVal: maxVal,
	}
}

type BitMap struct {
	data   []int64
	maxVal int // 最大值
}

func (obj *BitMap) Add(num int) {
	if num <= obj.maxVal {
		// obj.data[num / 64] |= (1 << (num % 64))
		obj.data[num>>6] |= (1 << (num & 63))
	}
}

func (obj *BitMap) Delete(num int) {
	if num <= obj.maxVal {
		obj.data[num>>6] &= ^(1 << (num & 63))
	}
}

func (obj *BitMap) Find(num int) bool {
	if num <= obj.maxVal {
		return obj.data[num>>6]&(1<<(num&63)) != 0
	}

	return false
}
