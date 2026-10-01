package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

// combinationSum1 返回所有可能的组合，使得数组 candidates 中的元素和等于 target
func combinationSum1(candidates []int, target int) [][]int {
	return process1(candidates, len(candidates)-1, target)
}

// process1 递归函数，返回在 candidates[0...index] 中选择元素，使得和为 target 的所有组合
func process1(arr []int, index int, target int) [][]int {
	ans := [][]int{}
	if target == 0 {
		ans = append(ans, []int{})
		return ans
	}
	if index == -1 {
		return ans
	}
	for zhang := 0; zhang*arr[index] <= target; zhang++ {
		preLists := process1(arr, index-1, target-(zhang*arr[index]))
		for _, pre := range preLists {
			for i := 0; i < zhang; i++ {
				pre = append(pre, arr[index])
			}
			ans = append(ans, pre)
		}
	}
	return ans
}

// combinationSum2 返回所有可能的组合，使得数组 candidates 中的元素和等于 target，使用记忆化搜索优化
func combinationSum2(candidates []int, target int) [][]int {
	mapIndexTarget := make(map[int]map[int][][]int)
	return process2(candidates, len(candidates)-1, target, mapIndexTarget)
}

// process2 递归函数，返回在 candidates[0...index] 中选择元素，使得和为 target 的所有组合，使用记忆化搜索优化
func process2(arr []int, index int, target int, mapIndexTarget map[int]map[int][][]int) [][]int {
	if val, ok := mapIndexTarget[index][target]; ok {
		return copyLists(val)
	}
	ans := [][]int{}
	if target == 0 {
		ans = append(ans, []int{})
		setAns(index, target, ans, mapIndexTarget)
		return copyLists(mapIndexTarget[index][target])
	}
	if index == -1 {
		setAns(index, target, ans, mapIndexTarget)
		return copyLists(mapIndexTarget[index][target])
	}
	for zhang := 0; zhang*arr[index] <= target; zhang++ {
		preLists := process2(arr, index-1, target-(zhang*arr[index]), mapIndexTarget)
		for _, pre := range preLists {
			for i := 0; i < zhang; i++ {
				pre = append(pre, arr[index])
			}
			ans = append(ans, pre)
		}
	}
	setAns(index, target, ans, mapIndexTarget)
	return copyLists(mapIndexTarget[index][target])
}

// setAns 将结果存储到 map 中
func setAns(index int, target int, ans [][]int, mapIndexTarget map[int]map[int][][]int) {
	if _, ok := mapIndexTarget[index]; !ok {
		mapIndexTarget[index] = make(map[int][][]int)
	}
	if _, ok := mapIndexTarget[index][target]; !ok {
		mapIndexTarget[index][target] = ans
	}
}

// copyLists 复制二维切片
func copyLists(lists [][]int) [][]int {
	ans := [][]int{}
	for _, cur := range lists {
		n := []int{}
		for _, num := range cur {
			n = append(n, num)
		}
		ans = append(ans, n)
	}
	return ans
}

// 最新版本
func combinationSum(candidates []int, target int) [][]int {
	path := make([]int, 0)
	ans := make([][]int, 0)
	process1(candidates, 0, target, path, &ans)
	return ans
}

func process1(arr []int, index int, rest int, path []int, ans *[][]int) {
	if index == len(arr) {
		if rest == 0 {
			// 收集答案
			tmp := make([]int, len(path))
			copy(tmp, path)
			*ans = append(*ans, tmp)
		}
		return
	}

	// 不要index位置的数
	process1(arr, index+1, rest, path, ans)

	// 要index位置的数
	for count := 1; count*arr[index] <= rest; count++ {
		// 思考这里为什么可以这么写？
		path = append(path, arr[index])
		process1(arr, index+1, rest-count*arr[index], path, ans)
	}
}
