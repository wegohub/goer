package base

import (
	"fmt"
	"sync"
	"time"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func mapReduceV1() {

	// 任务数量
	taskNum := 10

	// 完成任务标记队列
	ch := make(chan struct{}, taskNum)
	defer close(ch)

	// 开启任务
	for i := 0; i < taskNum; i++ {
		go func() {
			defer func() {
				ch <- struct{}{}
			}()
			<-time.After(time.Second)
		}()
	}

	// 等待任务完成
	for i := 0; i < taskNum; i++ {
		<-ch
	}

	fmt.Println("任务完成1")
}

func mapReduceV2() {
	// 任务数量
	taskNum := 10

	// 数据通道
	dataChan := make(chan interface{})

	// 子协程处理结果
	ans := make([]interface{}, 0, taskNum)

	// 读协程处理完的信号
	stopChan := make(chan struct{})

	// 开启读协程
	go func(dataChan chan interface{}, ans *[]interface{}) {
		for item := range dataChan {
			*ans = append(*ans, item)
		}
		stopChan <- struct{}{}
	}(dataChan, &ans)

	// 开启子任务
	var wg sync.WaitGroup
	for i := 0; i < taskNum; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			dataChan <- time.Now().Nanosecond()
		}()
	}

	// 阻塞等待子任务完成
	wg.Wait()
	// 关闭数据通道，读协程退出
	close(dataChan)
	<-stopChan

	// 这里可能读到的数据不完整
	fmt.Println(ans)
	fmt.Println("任务完成2")
}

// V2的优雅版
func mapReduceV3() {
	taskNum := 10

	// 数据通道
	dataChan := make(chan interface{})

	// 开启写任务
	go func(dataChan chan interface{}) {
		var wg sync.WaitGroup
		for i := 0; i < taskNum; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				dataChan <- time.Now().Nanosecond()
			}()
		}
		wg.Wait()
		close(dataChan)
	}(dataChan)

	ans := make([]interface{}, 0, taskNum)
	for item := range dataChan {
		ans = append(ans, item)
	}

	fmt.Println(ans)
	fmt.Println("任务完成3")
}
