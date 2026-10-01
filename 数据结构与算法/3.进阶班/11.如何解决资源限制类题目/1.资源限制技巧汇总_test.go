package class12

import "time"

// @timeout: 15
func main(params map[string]interface{}) (interface{}, error) {
	time.Sleep(10 * time.Second)
	return "Hello World!", nil
}
