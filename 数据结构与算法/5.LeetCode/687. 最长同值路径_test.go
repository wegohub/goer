package leetcode

import . "backend/utils/algo"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return process(root).Max - 1
}

type Info struct {
	// 在一条路径上：要求每个节点通过且只通过一遍
	Len int // 路径必须从x出发且只能往下走的情况下，路径的最大距离
	Max int // 路径不要求必须从x出发的情况下，整棵树的合法路径最大距离
}

func process(x *TreeNode) *Info {
	if x == nil {
		return &Info{Len: 0, Max: 0}
	}

	l := x.Left
	r := x.Right
	linfo := process(l)
	rinfo := process(r)

	len := 1
	if l != nil && l.Val == x.Val {
		len = linfo.Len + 1
	}
	if r != nil && r.Val == x.Val {
		len = Max(len, rinfo.Len+1)
	}

	max := Max(Max(linfo.Max, rinfo.Max), len)
	if l != nil && r != nil && l.Val == x.Val && r.Val == x.Val {
		max = Max(max, linfo.Len+rinfo.Len+1)
	}

	return &Info{Len: len, Max: max}
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
