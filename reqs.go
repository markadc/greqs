package main

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

type S = map[string]string
type A = map[string]any

// Response 响应
type Response struct {
	*http.Response
	Body []byte
}

// Text 响应的文本数据
func (r *Response) Text() string {
	return string(r.Body)
}

// JSON 响应的 JSON 数据
func (r *Response) JSON() (map[string]any, error) {
	var jsonMap map[string]any
	err := json.Unmarshal(r.Body, &jsonMap)
	if err != nil {
		return nil, err
	}
	return jsonMap, err
}

// JSONString 响应的 JSON 字符串
func (r *Response) JSONString() (string, error) {
	jsonMap, err := r.JSON()
	if err != nil {
		return "", err
	}
	byteArr, err := json.Marshal(jsonMap)
	if err != nil {
		return "", err
	}
	jsonStr := string(byteArr)
	return jsonStr, err
}

// PrettyJSONString 响应的 JSON 字符串（适合输出展示）
func (r *Response) PrettyJSONString() (string, error) {
	var buf bytes.Buffer
	err := json.Indent(&buf, r.Body, "", "  ")
	if err != nil {
		return "", fmt.Errorf("格式化JSON失败: %w", err)
	}
	return buf.String(), nil
}

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
func Do(client *http.Client, req *http.Request) (*Response, error) {
	resp, err := client.Do(req)
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

// PostForm 发送 POST 请求（表单形式）
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

// 请求配置
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

func SendGet(url string, opts *Options) (*Response, error) {
	return Send("GET", url, opts)
}

func SendPost(url string, opts *Options) (*Response, error) {
	return Send("POST", url, opts)
}
