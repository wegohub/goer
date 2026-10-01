package class15

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

type NumMatrix struct {
	tree [][]int // 树状数
	nums [][]int // 原始数据
	N    int     // 行数
	M    int     // 列数
}

func Constructor(matrix [][]int) NumMatrix {
	N := len(matrix)
	M := len(matrix[0])
	tree := make([][]int, N+1)
	for i := 0; i < N+1; i++ {
		tree[i] = make([]int, M+1)
	}
	nums := make([][]int, N)
	for i := 0; i < N; i++ {
		nums[i] = make([]int, M)
	}
	indexTree := NumMatrix{
		tree: tree,
		nums: nums,
		N:    N,
		M:    M,
	}
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			indexTree.Update(i, j, matrix[i][j])
		}
	}
	return indexTree
}

// [row, col]位置的值更新到val
func (this *NumMatrix) Update(row int, col int, val int) {
	if this.N == 0 || this.M == 0 {
		return
	}
	add := val - this.nums[row][col]
	this.nums[row][col] = val
	for i := row + 1; i <= this.N; i += i & (-i) {
		for j := col + 1; j <= this.M; j += j & (-j) {
			this.tree[i][j] += add
		}
	}
}

// 返回 [0,0] 点 到 [row, col] 区域的累加和
func (this *NumMatrix) sum(row, col int) int {
	sum := 0
	for i := row + 1; i > 0; i -= i & (-i) {
		for j := col + 1; j > 0; j -= j & (-j) {
			sum += this.tree[i][j]
		}
	}
	return sum
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	if this.N == 0 || this.M == 0 {
		return 0
	}
	return this.sum(row2, col2) - this.sum(row2, col1-1) - this.sum(row1-1, col2) + this.sum(row1-1, col1-1)
}
