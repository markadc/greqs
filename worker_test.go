package greqs

import (
	"fmt"
	"net/http"
	_url "net/url"
	"testing"
	"time"
)

// 请求中间件，可以补充请求头
func m1(req *http.Request) {
	req.Header.Set("User-Agent", "Greqs/0.0.1")
}

// 客户端中间件，可以设置代理
func m2(cli *http.Client) {
	proxyURL, _ := _url.Parse("http://127.0.0.1:7890")
	cli.Transport = &http.Transport{Proxy: http.ProxyURL(proxyURL)}
}

func TestWorker_Get(t *testing.T) {
	worker := NewWorker("", 5*time.Second, m1, m2)

	url := "https://httpbin.org/get"
	resp, err := worker.Get(url, nil)
	if err != nil {
		t.Errorf("worker.Get err: %v", err)
	}

	fmt.Println(resp.Text())
	fmt.Printf("%+v\n", resp.Request.Header)
	fmt.Println(resp.Request.Header.Get("Name"))
}
