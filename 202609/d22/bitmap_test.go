package d22

import (
	"fmt"
	"math"
	"math/rand/v2"
	"strings"
	"testing"
)

// PrinBinary 打印一个int64数字的二进制位
func PrinBinary(num int64) {
	sb := strings.Builder{}
	for i := 63; i >= 0; i-- {
		bit := "0"
		if (num & (int64(1) << i)) != 0 {
			bit = "1"
		}
		sb.WriteString(bit)
	}
	fmt.Println(sb.String())
}

// Test_BitAddDeleteGet BitMap的增、删、查
func Test_BitMapAddDeleteGet(t *testing.T) {

	// 一个 int64 8字节能表示的数字是 0-63
	var bitMap int64

	// 新增一个数 32
	num := 32
	bitMap = bitMap | (int64(1) << num)
	PrinBinary(bitMap)

	// 查询num是否存在
	for i := 0; i < 64; i++ {
		isExist := bitMap&(int64(1)<<i) != 0
		fmt.Printf("%d是否存在: %v\n", i, isExist)
	}

	// 删除num
	bitMap = bitMap & ^(int64(1) << num)
	PrinBinary(bitMap)
}

// BitMap 位图实现
type BitMap struct {
	bits []uint64
	max  int64
}

// NewBitMap 输入最大值max, 在位图上能表示 [-max, +max] 的数
func NewBitMap(max int64) *BitMap {
	if max < 0 {
		panic("the maximum value must be greater than or equal to 0")
	}

	// 确保 2max + 1 + 63 不会溢出
	if max > (math.MaxInt64-63-1)/2 {
		panic("max is too large")
	}

	// 计算要能表示 [-max, max] 所需要的比特位
	// 如果只表示 [0, max], 即不需要表示负数, bitCount := max + 1
	// 0 0
	// 100 100
	// -100 0
	// 199 1
	bitCount := 2*max + 1

	// 计算所需的word数量，像上取整数，需要每组 N 个元素时：ceil(x / N) = (x + N - 1) / N
	// eg: bitCount = 64
	// (64 + 63) / 64 = 127 / 64 = 1 只需要一个 uint64
	wordCount := (bitCount + 63) >> 6

	return &BitMap{
		bits: make([]uint64, wordCount),
		max:  max,
	}
}

// position 计算一个数字在bits中数组的下标和位数
func (obj *BitMap) position(num int64) (wordIndex int64, bitIndex uint64, err error) {
	if num < ^obj.max+1 || num > obj.max {
		return 0, 0, fmt.Errorf("num must in [%d, %d]", -obj.max, obj.max)
	}

	// 计算偏移量 100
	// -100 + 100 = 0
	// -99 + 100 = 1
	// 0 + 100 = 100
	// 1 + 100 = 101
	// 100 + 100 = 200
	// [0, 100]
	// 0 = 0
	// 1= 1
	offset := num + obj.max

	// 计算wordIndex
	// [int64, int64]
	wordIndex = offset >> 6

	// 计算bitIndex
	bitIndex = uint64(offset & 63)

	return
}

// Add 向BitMap中添加一个数字
func (obj *BitMap) Add(num int64) error {
	wordIndex, bitIndex, err := obj.position(num)
	if err != nil {
		return err
	}

	obj.bits[wordIndex] |= uint64(1) << bitIndex

	return nil
}

// Delete 向BitMap中删除一个数字
func (obj *BitMap) Delete(num int64) error {
	wordIndex, bitIndex, err := obj.position(num)
	if err != nil {
		return err
	}

	// Go中按位清零常常这样写: obj.bits[wordIndex] &^= uint64(1) << bitIndex
	obj.bits[wordIndex] &= ^(uint64(1) << bitIndex)

	return nil
}

// Contains BitMap中是否包含一个数字
func (obj *BitMap) Contains(num int64) (bool, error) {
	wordIndex, bitIndex, err := obj.position(num)
	if err != nil {
		return false, err
	}

	isExist := obj.bits[wordIndex]&(uint64(1)<<bitIndex) != 0

	return isExist, nil
}

// 功能测试
func Test_BitMap_Feature(t *testing.T) {
	bm := NewBitMap(100)

	var num int64 = 50

	_ = bm.Add(num)

	isExist, _ := bm.Contains(num)
	fmt.Printf("数字num: %d 是否存在: %v\n", num, isExist)

	_ = bm.Delete(num)

	isExistAfterDelete, _ := bm.Contains(num)
	fmt.Printf("数字num: %d 是否存在: %v\n", num, isExistAfterDelete)
}

// HashSet 作为BitMap的对数器使用
type HashSet struct {
	data map[int64]struct{}
	max  int64
}

// NewHashSet 实例化
func NewHashSet(max int64) *HashSet {
	return &HashSet{
		data: make(map[int64]struct{}),
		max:  max,
	}
}

// Add 添加一个数字
func (obj *HashSet) Add(num int64) error {
	if num < -obj.max || num > obj.max {
		return fmt.Errorf("num must in [%d, %d]", -obj.max, obj.max)
	}

	obj.data[num] = struct{}{}

	return nil
}

// Delete 删除一个数字
func (obj *HashSet) Delete(num int64) error {
	if num < -obj.max || num > obj.max {
		return fmt.Errorf("num must in [%d, %d]", -obj.max, obj.max)
	}

	delete(obj.data, num)

	return nil
}

// Contains 包含一个数字
func (obj *HashSet) Contains(num int64) (bool, error) {
	if num < -obj.max || num > obj.max {
		return false, fmt.Errorf("num must in [%d, %d]", -obj.max, obj.max)
	}

	if _, ok := obj.data[num]; ok {
		return true, nil
	}

	return false, nil
}

// 验证BitMap实现的正确性
// 测试范围：[0, max] max=10000
// 测试次数：10000000次
// 随机操作：1/3概率添加，1/3概率删除，1/3概率查询
func Test_BitMapLogarithm(t *testing.T) {
	var (
		testNum       = 10000000
		maxNum  int64 = 10000
	)

	bm := NewBitMap(maxNum)
	hs := NewHashSet(maxNum)

	for i := 0; i < testNum; i++ {
		num := rand.Int64N(2*maxNum+1) - maxNum // [-maxNum, +maxNum]
		switch i % 3 {
		case 0:
			err := bm.Add(num)
			if err != nil {
				t.Fatal(err)
			}
			err = hs.Add(num)
			if err != nil {
				t.Fatal(err)
			}
		case 1:
			err := bm.Delete(num)
			if err != nil {
				t.Fatal(err)
			}
			err = hs.Delete(num)
			if err != nil {
				t.Fatal(err)
			}
		case 2:
			ans1, err := bm.Contains(num)
			if err != nil {
				t.Fatal(err)
			}
			ans2, err := hs.Contains(num)
			if err != nil {
				t.Fatal(err)
			}
			if ans1 != ans2 {
				t.Fatal("not equal")
			}
		}

	}

	// 最终结果校验
	for i := -maxNum; i <= maxNum; i++ {
		ans1, _ := bm.Contains(i)
		ans2, _ := hs.Contains(i)
		if ans1 != ans2 {
			t.Fatal("not equal finally")
		}
	}
}
