package base

import (
	"fmt"
	"sync"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	syncMap()
	return "Hello World!", nil
}

func syncMap() {
	var mp sync.Map
	mp.Store("key1", "val1")
	mp.Store("key2", "val2")
	v, ok := mp.Load("key1")
	if ok {
		fmt.Println(v)
	} else {
		panic("key1 not exist")
	}

	mp.Delete("key1")

	// 熔断器
	mp.Range(func(key, value any) bool {
		fmt.Println(key, value)
		return true
	})

}
