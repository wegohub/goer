package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func numIslands(grid [][]byte) int {
	n := len(grid)
	if n == 0 || len(grid[0]) == 0 {
		return 0
	}
	set := NewUnionFind2(grid)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				// 跟上面的合并
				if i-1 >= 0 && grid[i-1][j] == '1' {
					set.Union(i, j, i-1, j)
				}
				// 跟左边合并
				if j-1 >= 0 && grid[i][j-1] == '1' {
					set.Union(i, j, i, j-1)
				}
			}
		}
	}
	return set.GetSets()
}

func NewUnionFind2(grid [][]byte) *UnionFind2 {
	row := len(grid)
	col := len(grid[0])
	length := row * col
	parent := make([]int, length)
	size := make([]int, length)
	help := make([]int, length)
	sets := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				index := i*col + j
				parent[index] = index
				size[index] = 1
				sets++
			}
		}
	}
	return &UnionFind2{
		parent: parent,
		size:   size,
		help:   help,
		col:    col,
		sets:   sets,
	}
}

type UnionFind2 struct {
	parent []int
	size   []int
	help   []int
	col    int // 总列数
	sets   int // 岛屿数
}

func (obj *UnionFind2) index(i, j int) int {
	return i*obj.col + j
}

func (obj *UnionFind2) Union(i1, j1, i2, j2 int) {
	i := obj.index(i1, j1)
	j := obj.index(i2, j2)
	iHead := obj.find(i)
	jHead := obj.find(j)
	if iHead != jHead {
		if obj.size[iHead] > obj.size[jHead] {
			obj.parent[jHead] = iHead
			obj.size[iHead] += obj.size[jHead]
		} else {
			obj.parent[iHead] = jHead
			obj.size[jHead] += obj.size[iHead]
		}
		obj.sets--
	}
}

func (obj *UnionFind2) find(cur int) int {
	height := 0
	for cur != obj.parent[cur] {
		obj.help[height] = cur
		height++
		cur = obj.parent[cur]
	}
	for height--; height >= 0; height-- {
		obj.parent[obj.help[height]] = cur
	}
	return cur
}

func (obj *UnionFind2) GetSets() int {
	return obj.sets
}

// 感染法 O(M*N)
func numIslands1(grid [][]byte) int {
	n := len(grid)
	if n == 0 || len(grid[0]) == 0 {
		return 0
	}
	num := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				num++
				infection(grid, i, j)
			}
		}
	}
	return num
}

// 从i,j开始感染上下左右的1
func infection(gird [][]byte, i, j int) {
	if i < 0 || i >= len(gird) || j < 0 || j >= len(gird[0]) || gird[i][j] != '1' {
		return
	}
	gird[i][j] = '0'
	infection(gird, i-1, j)
	infection(gird, i+1, j)
	infection(gird, i, j-1)
	infection(gird, i, j+1)
}
