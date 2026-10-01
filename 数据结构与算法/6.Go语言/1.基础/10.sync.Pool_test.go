package base

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

type Student struct {
	Name     string
	Age      int
	CreateAt int64
}

func NewStudent() *Student {
	return &Student{
		Name:     "",
		Age:      0,
		CreateAt: time.Now().UnixNano(),
	}
}

func Test_SyncPool() {
	pool := sync.Pool{}

	// 池子中没有对象时调用的构造方法
	pool.New = func() any {
		return NewStudent()
	}

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			obj := pool.Get().(*Student)
			fmt.Println(obj.CreateAt)
			<-time.After(time.Second * time.Duration(rand.Intn(3)))
			pool.Put(obj)
		}()
	}
	wg.Wait()
}
