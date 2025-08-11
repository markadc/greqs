package main

import (
	"fmt"
	"greqs/requests"
)

func main() {
	worker := requests.NewWorker()

	formData := map[string]string{"name": "form"}
	opts := &requests.Options{FormData: formData}

	resp, _ := worker.Post("https://httpbin.org/post", opts)

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Text())
}
