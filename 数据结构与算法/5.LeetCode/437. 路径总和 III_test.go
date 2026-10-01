package leetcode

import . "backend/utils/algo"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

// 先序遍历
func pathSum(root *TreeNode, targetSum int) int {
	mp := make(map[int]int)
	mp[0] = 1
	return process(root, targetSum, 0, mp)
}

func process(x *TreeNode, sum int, preAll int, mp map[int]int) int {
	if x == nil {
		return 0
	}
	all := preAll + x.Val
	ans := 0
	if v, ok := mp[all-sum]; ok {
		ans = v
	}
	if v, ok := mp[all]; ok {
		mp[all] = v + 1
	} else {
		mp[all] = 1
	}

	ans += process(x.Left, sum, all, mp)
	ans += process(x.Right, sum, all, mp)
	if mp[all] == 1 {
		delete(mp, all)
	} else {
		mp[all] = mp[all] - 1
	}

	return ans
}
