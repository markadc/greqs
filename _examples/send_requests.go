package main

import (
	"encoding/json"
	"fmt"
	"greqs/log"
	"greqs/requests"
	"time"
)

var (
	worker  = requests.NewWorker()
	p       = log.NewPrinter()
	headers = map[string]string{"User-Agent": "Wauo"}
)

func main() {
	fmt.Println("====== GET ======")
	UseGet()

	fmt.Println("====== POST FormData ======")
	UsePost1()

	fmt.Println("====== POST JSON ======")
	UsePost2()

	fmt.Println("====== Proxy1 ======")
	UseProxy1()

	fmt.Println("====== Proxy2 ======")
	UseProxy2()

}

func UseGet() {
	params := map[string]string{"page": "1", "limit": "10"}
	opts := &requests.Options{Params: params, Headers: headers}
	resp, _ := worker.Get("https://httpbin.org/get", opts)
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Text())
}

func UsePost1() {
	formData := map[string]string{"name": "form"}
	opts := &requests.Options{FormData: formData, Headers: headers}
	resp, _ := worker.Post("https://httpbin.org/post", opts)
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Text())
}

func UsePost2() {
	jsonMap := map[string]string{"name": "json"}
	b, _ := json.Marshal(jsonMap)
	opts := &requests.Options{JSON: b, Headers: headers}
	resp, _ := worker.Post("https://httpbin.org/post", opts)
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Text())

	p.Red("Call...JSON")
	j, _ := resp.JSON()
	fmt.Println(j)

	p.Red("Call...JSONString")
	js, _ := resp.JSONString()
	fmt.Println(js)

	p.Red("Call...PrettyJSONString")
	pjs, _ := resp.PrettyJSONString()
	fmt.Println(pjs)
}

func UseProxy1() {
	opts := &requests.Options{
		Proxy:   "http://127.0.0.1:7890", // 具体的代理
		Headers: headers,
	}

	resp, err := worker.Get("https://httpbin.org/ip", opts)
	if err != nil {
		fmt.Printf("Error => %s\n", err)
		return
	}

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Text())
}
func UseProxy2() {
	worker.SetDefaultProxy("http://127.0.0.1:7890")
	worker.SetDefaultTimeout(3 * time.Second)

	resp, err := worker.Get("https://httpbin.org/ip", nil)
	if err != nil {
		fmt.Printf("Error => %s\n", err)
		return
	}

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Text())
}
