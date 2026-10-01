package class03

import (
	"fmt"
	"math/rand"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	// 模拟全球登录用户通道
	users := make(chan string)

	go func() {
		for i := 1; i <= 1000000; i++ {
			users <- fmt.Sprintf("user%d", i)
		}
		close(users)
	}()

	// 抽取100个幸运观众
	luckyUsers := ReservoirSampling(users, 100)

	fmt.Println("幸运观众名单：")
	for _, user := range luckyUsers {
		fmt.Println(user)
	}
	return "Hello World!", nil
}

// ReservoirSampling 使用蓄水池抽样算法从全球登录用户中抽取100个幸运观众
func ReservoirSampling(users chan string, k int) []string {
	// 初始化蓄水池
	reservoir := make([]string, k)
	i := 0

	// 遍历用户通道
	for user := range users {
		if i < k {
			// 填充蓄水池的前k个元素
			reservoir[i] = user
		} else {
			// 以k/i的概率替换蓄水池中的元素
			j := rand.Intn(i + 1)
			if j < k {
				reservoir[j] = user
			}
		}
		i++
	}

	return reservoir
}
