package base

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func server() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})
	http.ListenAndServe(":8085", nil)
}

func client() {
	reqBody, _ := json.Marshal(map[string]string{"key1": "val1", "key2": "val2"})

	resp, _ := http.Post(":8085", "application/json", bytes.NewReader(reqBody))
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	fmt.Printf("resp: %s", respBody)
}
