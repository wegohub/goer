package base

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

// 1. context.WithCancel(context.TODO()) 创建newCancelCtx(parent)， 并启动协程监听父级ctx的Done信号， 返回cancelCtx和取消闭包，调用方法将执行c.cancel(true, Canceled)
// 将自己从父节点的children中移除，并移除children集合中的的所以子ctx，并close(done)通道
// 2. 对于后续的子协程监听父级ctx.Done()信号做出退出判断
// 3. 后续的子协程创建cancelCtx会检查父级是不是cancelCtx，如果是则直接加入父级ctx的children集合，不是则开启协程监听父级的done和自身的done信号
func Test_CancelContext(t *testing.T) {
	// 创建一个通道来接收信号
	sigChan := make(chan os.Signal, 1)
	// 注册要监听的信号
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.TODO())
	for i := 0; i < 10; i++ {
		go func(ctx context.Context, idx int) {
		loop:
			for {
				select {
				case <-ctx.Done():
					fmt.Println(fmt.Sprintf("子协程%d退出: %s", idx, ctx.Err().Error()))
					break loop
				default:
					fmt.Printf("子协程%d业务处理\n", idx)
					time.Sleep(time.Second)
				}
			}
		}(ctx, i)
	}

	go func() {
		sig := <-sigChan
		fmt.Printf("接收到信号: %v\n", sig)
		cancel()
	}()

	select {
	case <-ctx.Done():
		time.Sleep(time.Second)
		fmt.Println("主流程退出")
	}
}

// 无缓存通道
func Test_Channel(t *testing.T) {
	c := make(chan struct{})

	go func(c chan struct{}) {
		<-c
		fmt.Println("协程1退出")
	}(c)

	go func(c chan struct{}) {
		<-c
		fmt.Println("协程2退出")
	}(c)

	time.Sleep(time.Second)
	// c <- struct{}{} // 这样只会有一个协程能监听到信号
	close(c) // 这样两个协程都会收到信号
	time.Sleep(time.Second)
}

func Test_TimerContext(t *testing.T) {
	// 创建一个通道来接收信号
	sigChan := make(chan os.Signal, 1)
	// 注册要监听的信号
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
	for i := 0; i < 10; i++ {
		go func(ctx context.Context, idx int) {
		loop:
			for {
				select {
				case <-ctx.Done():
					fmt.Println(fmt.Sprintf("子协程%d退出: %s", idx, ctx.Err().Error()))
					break loop // 或者return
				default:
					fmt.Printf("子协程%d业务处理\n", idx)
					time.Sleep(time.Second)
				}
			}
		}(ctx, i)
	}

	go func() {
		sig := <-sigChan
		fmt.Printf("接收到信号: %v\n", sig)
		cancel()
	}()

	select {
	case <-ctx.Done():
		time.Sleep(time.Second)
		fmt.Println("主流程退出", ctx.Err())
	}
}
