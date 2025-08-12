package greqs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	_url "net/url"
	"strings"
	"time"
)

// MakeUrl 构建完整的 url 地址
func MakeUrl(url string, params S) string {
	q := _url.Values{}
	for key, val := range params {
		q.Add(key, val)
	}
	if !strings.Contains(url, "?") {
		url += "?"
	}
	url += q.Encode()
	return url
}

// SetHeaders 为请求设置请求头
func SetHeaders(req *http.Request, headers S) {
	for key, val := range headers {
		req.Header.Set(key, val)
	}
}

// MakeGetRequest 创建 GET 请求
func MakeGetRequest(url string, headers S) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	SetHeaders(req, headers)
	return req, nil
}

// MakePostRequest 创建 POST 请求（JSON 形式）
func MakePostRequest(url string, headers S, data A) (*http.Request, error) {
	body, _ := json.Marshal(data)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	SetHeaders(req, headers)
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

// MakePostFormRequest 创建 POST 请求（表单形式）
func MakePostFormRequest(url string, headers S, form S) (*http.Request, error) {
	val := _url.Values{}
	for k, v := range form {
		val.Set(k, v)
	}
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(val.Encode()))
	if err != nil {
		return nil, err
	}
	SetHeaders(req, headers)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return req, nil
}

// GetClient 获取客户端
func GetClient(proxy string, timeout time.Duration) *http.Client {
	client := &http.Client{}
	if proxy != "" {
		proxyURL, err := _url.Parse(proxy)
		if err != nil {
			panic(err)
		}
		client.Transport = &http.Transport{Proxy: http.ProxyURL(proxyURL)}
	}
	if timeout > 0 {
		client.Timeout = timeout
	}
	return client
}

// Do 发送请求，获取响应
func Do(cli *http.Client, req *http.Request) (*Response, error) {
	resp, err := cli.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return &Response{resp, bodyBytes}, nil
}

func RandInt(min, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	num := min + r.Intn(max-min+1)
	return num
}

// Options 请求配置
type Options struct {
	Params  S
	Headers S
	Data    A
	Form    S
	Proxy   string
	Timeout time.Duration
}

// Send 发送请求
func Send(method, url string, opts *Options) (*Response, error) {
	method = strings.ToUpper(method)
	if method != "GET" && method != "POST" {
		return nil, fmt.Errorf("仅支持 GET、POST，不支持 %s", method)
	}

	if opts == nil {
		opts = &Options{}
	}
	cli := GetClient(opts.Proxy, opts.Timeout)

	if method == "GET" {
		if opts.Params != nil {
			url = MakeUrl(url, opts.Params)
		}
		req, err := MakeGetRequest(url, opts.Headers)
		if err != nil {
			return nil, err
		}
		return Do(cli, req)

	} else {
		if opts.Data != nil {
			req, err := MakePostRequest(url, opts.Headers, opts.Data)
			if err != nil {
				return nil, err
			}
			return Do(cli, req)

		} else if opts.Form != nil {
			req, err := MakePostFormRequest(url, opts.Headers, opts.Form)
			if err != nil {
				return nil, err
			}
			return Do(cli, req)
		} else {
			return nil, fmt.Errorf("无效的 POST 请求")
		}
	}
}

// SendGetRequest 发送 GET 请求
func SendGetRequest(url string, opts *Options) (*Response, error) {
	return Send("GET", url, opts)
}

// SendPostRequest 发送 POST 请求（根据 opts 参数自动识别 JSON 或者 Form 形式）
func SendPostRequest(url string, opts *Options) (*Response, error) {
	return Send("POST", url, opts)
}
