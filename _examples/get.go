package main

import (
	"fmt"
	"greqs/requests"
)

func main() {
	worker := requests.NewWorker()

	params := map[string]string{"page": "1", "limit": "10"}
	headers := map[string]string{"User-Agent": "Greqs"}
	opts := &requests.Options{Params: params, Headers: headers}

	resp, _ := worker.Get("https://httpbin.org/get", opts)

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Text())
}
