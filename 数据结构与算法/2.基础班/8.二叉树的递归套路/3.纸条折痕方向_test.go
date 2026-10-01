package class08

import "fmt"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	printProcess(1, 2, true)
	return "Hello World!", nil
}

func printProcess(i int, N int, isDown bool) {
	if i > N {
		return
	}

	printProcess(i+1, N, true)
	if isDown {
		fmt.Println("凹")
	} else {
		fmt.Println("凸")
	}
	printProcess(i+1, N, false)
}
