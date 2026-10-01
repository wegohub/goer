package leetcode

import . "backend/utils/algo"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return process(root, p, q).Ans
}

type Info struct {
	Ans   *TreeNode
	FindP bool
	FindQ bool
}

func process(root, p, q *TreeNode) *Info {
	if root == nil {
		return &Info{}
	}

	left := process(root.Left, p, q)
	right := process(root.Right, p, q)

	findP := root == p || left.FindP || right.FindP
	findQ := root == q || left.FindQ || right.FindQ

	var ans *TreeNode
	if left.Ans != nil {
		ans = left.Ans
	}
	if right.Ans != nil {
		ans = right.Ans
	}
	if ans == nil {
		if findP && findQ {
			ans = root
		}
	}

	return &Info{
		Ans:   ans,
		FindP: findP,
		FindQ: findQ,
	}
}
