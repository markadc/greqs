package main

import (
	"encoding/json"
	"fmt"
	"greqs/requests"
)

func main() {
	worker := requests.NewWorker()

	jsonMap := map[string]string{"name": "json"}
	b, _ := json.Marshal(jsonMap)
	opts := &requests.Options{JSON: b}

	resp, _ := worker.Post("https://httpbin.org/post", opts)

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Text())
}
