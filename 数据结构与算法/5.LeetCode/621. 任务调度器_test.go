package leetcode

import "math"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func leastInterval(tasks []byte, free int) int {
	count := make([]int, 256)
	// 出现最多次的任务，到底是出现了几次
	maxCount := 0
	for _, task := range tasks {
		count[task]++
		maxCount = int(math.Max(float64(maxCount), float64(count[task])))
	}
	// 有多少种任务，都出现最多次
	maxKinds := 0
	for _, cnt := range count {
		if cnt == maxCount {
			maxKinds++
		}
	}
	// 砍掉最后一组剩余的任务数
	tasksExceptFinalTeam := len(tasks) - maxKinds
	spaces := (free + 1) * (maxCount - 1)
	restSpaces := int(math.Max(0, float64(spaces-tasksExceptFinalTeam)))
	return len(tasks) + restSpaces
}
