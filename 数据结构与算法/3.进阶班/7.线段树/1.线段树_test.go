package class07

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func NewSegmentTree(origin []int) *SegmentTree {
	maxN := len(origin) + 1
	// 原始序列到线段树序列的映射，0位置弃而不用
	// 结论：任何位置i的左右孩子下标是: 2*i  2*i+1
	// 任何位置i的父级节点是: i/2
	arr := make([]int, maxN)
	for i, v := range origin {
		arr[i+1] = v
	}
	// 用数组模拟线段树，4N长度一定装得下
	return &SegmentTree{
		maxN:     maxN,
		arr:      arr,
		sum:      make([]int, maxN<<2),
		lazy:     make([]int, maxN<<2),
		change:   make([]int, maxN<<2),
		isUpdate: make([]bool, maxN<<2),
	}
}

// SegmentTree 线段树
type SegmentTree struct {
	maxN int   // 原始数据的长度(从下标1开始，0位置不使用)
	arr  []int // 原始序列映射为从1开始的的
	// 累加信息
	sum  []int // 维护线段树累加区间和
	lazy []int // 累加和懒惰标记
	// 更新信息
	change   []int  // 维护线段树更新区间
	isUpdate []bool // 更新慵懒标记
}

// pushUp 在rt位置向左右孩子收集信息
func (obj *SegmentTree) pushUp(rt int) {
	obj.sum[rt] = obj.sum[rt<<1] + obj.sum[rt<<1|1]
}

// pushDown 在rt位置向左右孩子下发懒加信息
func (obj *SegmentTree) pushDown(rt, ln, rn int) {
	// 下发更新操作
	if obj.isUpdate[rt] {
		// 左右孩子标记为更新状态
		obj.isUpdate[rt<<1] = true
		obj.isUpdate[rt<<1|1] = true
		// 左右还是标记更新值
		obj.change[rt<<1] = obj.change[rt]
		obj.change[rt<<1|1] = obj.change[rt]
		// 将左右还是拦住的累加信息清空
		obj.lazy[rt<<1] = 0
		obj.lazy[rt<<1|1] = 0
		// 更新左右还是的累加信息
		obj.sum[rt<<1] = obj.change[rt] * ln
		obj.sum[rt<<1|1] = obj.change[rt] * rn
		// 标记rt位置下发更新完成
		obj.isUpdate[rt] = false
	}
	// 下发累加操作
	if obj.lazy[rt] != 0 {
		// 左右还是继承父级的累加信息
		obj.lazy[rt<<1] += obj.lazy[rt]
		obj.sum[rt<<1] += obj.lazy[rt] * ln
		obj.lazy[rt<<1|1] += obj.lazy[rt]
		obj.sum[rt<<1|1] += obj.lazy[rt] * rn
		// 下发累加信息完成
		obj.lazy[rt] = 0
	}
}

// Build 在初始阶段把sum信息填好
// 在arr[l~r]范围上，去build，1~N，
// rt : 这个范围在sum中的下标
func (obj *SegmentTree) Build(l, r, rt int) {
	if l == r {
		obj.sum[rt] = obj.arr[l]
		return
	}
	mid := (l + r) >> 1
	obj.Build(l, mid, rt<<1)     // 左孩子
	obj.Build(mid+1, r, rt<<1|1) // 右孩子
	obj.pushUp(rt)               // 收集父级信息
}

// Add 累加
// L R C 表示任务，在L...R上累加C
// l r rt 表示在线段l...r上执行任务，当前位置为rt
func (obj *SegmentTree) Add(L, R, C, l, r, rt int) {
	// 如果任务把此时的范围全包含了
	if L <= l && R >= r {
		obj.sum[rt] += C * (r - l + 1)
		obj.lazy[rt] += C
		return
	}
	// 任务没有把此时的范围包住
	mid := (l + r) >> 1
	// 把当前位置的信息下发
	obj.pushDown(rt, mid-l+1, r-mid)
	// mid在L...R范围上
	if L <= mid {
		obj.Add(L, R, C, l, mid, rt<<1)
	}
	if R > mid {
		obj.Add(L, R, C, mid+1, r, rt<<1|1)
	}
	// 向左右孩子收集父级信息
	obj.pushUp(rt)
}

// Update 更新
// L R C 表示任务，在L...R上所有的值变成C
// l r rt 表示在线段l...r上执行任务，当前位置为rt
func (obj *SegmentTree) Update(L, R, C, l, r, rt int) {
	// 任务把此时的范围包住了
	if L <= l && R >= r {
		obj.isUpdate[rt] = true
		obj.change[rt] = C
		obj.sum[rt] = C * (r - l + 1)
		obj.lazy[rt] = 0
		return
	}
	// 当前任务躲不掉，无法懒更新，要往下发
	mid := (l + r) >> 1
	obj.pushDown(rt, mid-l+1, r-mid)
	// 在L...R范围上
	if L <= mid {
		obj.Update(L, R, C, l, mid, rt<<1)
	}
	if R > mid {
		obj.Update(L, R, C, mid+1, r, rt<<1|1)
	}
	// 向左右孩子收集父级信息
	obj.pushUp(rt)
}

// Query 查询 1~6 累加和是多少？ 1~8 rt
func (obj *SegmentTree) Query(L, R, l, r, rt int) int {
	// 如果范围被包含
	if L <= l && R >= r {
		return obj.sum[rt]
	}
	mid := (l + r) >> 1
	obj.pushDown(rt, mid-l+1, r-mid)
	ans := 0
	if L <= mid {
		ans += obj.Query(L, R, l, mid, rt<<1)
	}
	if R > mid {
		ans += obj.Query(L, R, mid+1, r, rt<<1|1)
	}
	return ans
}

func NewSegmentRight(origin []int) *SegmentRight {
	arr := make([]int, len(origin)+1)
	for i, v := range origin {
		arr[i+1] = v
	}
	return &SegmentRight{
		arr: arr,
	}
}

// SegmentRight 线段树对数器
type SegmentRight struct {
	arr []int
}

func (obj *SegmentRight) Add(L, R, C int) {
	for i := L; i <= R; i++ {
		obj.arr[i] += C
	}
}

func (obj *SegmentRight) Update(L, R, C int) {
	for i := L; i <= R; i++ {
		obj.arr[i] = C
	}
}

func (obj *SegmentRight) Query(L, R int) int {
	ans := 0
	for i := L; i <= R; i++ {
		ans += obj.arr[i]
	}
	return ans
}
