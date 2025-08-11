package main

import (
	"fmt"
	"greqs/requests"
)

func main() {
	worker := requests.NewWorker()

	opts := &requests.Options{Proxy: "http://127.0.0.1:7890"}

	resp, err := worker.Get("https://httpbin.org/ip", opts)
	if err != nil {
		fmt.Printf("Error => %s\n", err)
		return
	}

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Text())
}
