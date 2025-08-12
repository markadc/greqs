package greqs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	_url "net/url"
	"strings"
	"time"
)

// SetHeaders 为请求设置请求头
func SetHeaders(req *http.Request, headers S) {
	for key, val := range headers {
		req.Header.Set(key, val)
	}
}

// DoRequest 发送请求，获取响应
func DoRequest(req *http.Request) (*Response, error) {
	// 发送请求
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("发送请求失败: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	// 读取响应
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %w", err)
	}

	// 返回响应结构体
	return &Response{resp, bodyBytes}, nil
}

// GetClient 获取一个客户端
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

// Get 发送 GET 请求
func Get(url string, headers S) (*Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	SetHeaders(req, headers)
	return DoRequest(req)

}

// POST 发送 POST 请求（JOSN 形式）
func Post(url string, data A, headers S) (*Response, error) {
	body, _ := json.Marshal(data)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	SetHeaders(req, headers)
	req.Header.Set("Content-Type", "application/json")
	return DoRequest(req)
}

// PostForm 发送 POST 请求（Form 形式）
func PostForm(url string, form S, headers S) (*Response, error) {
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
	return DoRequest(req)
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
