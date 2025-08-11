package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// 请求的一些配置
type Options struct {
	Params   map[string]string // URL查询参数
	Headers  map[string]string // 请求头
	JSON     []byte            // JSON请求体
	FormData map[string]string // 表单数据
	Timeout  time.Duration     // 请求超时时间
	Proxy    string            // 代理URL
}

// 响应对象
type Response struct {
	*http.Response
	Body []byte
}

// 工作者
type Worker struct {
	client *http.Client
}

// new 一个工作者
func NewWorker() *Worker {
	return &Worker{
		client: &http.Client{},
	}
}

// 设置默认超时
func (w *Worker) SetDefaultTimeout(timeout time.Duration) {
	w.client.Timeout = timeout
}

// 设置默认代理
func (w *Worker) SetDefaultProxy(proxy string) {
	proxyURL, err := url.Parse(proxy)
	if err != nil {
		panic(err)
	}
	w.client.Transport = &http.Transport{Proxy: http.ProxyURL(proxyURL)}
}

// 发送 GET 请求
func (c *Worker) Get(urlStr string, opts *Options) (*Response, error) {
	return c.Send("GET", urlStr, opts)
}

// 发送 POST 请求
func (c *Worker) Post(urlStr string, opts *Options) (*Response, error) {
	return c.Send("POST", urlStr, opts)
}

// 构造完整的 url 地址
func MakeUrl(urlStr string, params map[string]string) string {
	q := url.Values{}
	for key, val := range params {
		q.Add(key, val)
	}
	if strings.Contains(urlStr, "?") {
		urlStr += "&" + q.Encode()
	} else {
		urlStr += "?" + q.Encode()
	}
	return urlStr
}

// 发送 HTTP 请求，根据方法类型 ( GET / POST ) 分别处理
func (c *Worker) Send(method, urlStr string, opts *Options) (*Response, error) {
	method = strings.ToUpper(method)
	if method != "GET" && method != "POST" {
		return nil, fmt.Errorf("不支持的HTTP方法: %s", method)
	}

	// 处理查询字符串
	if opts != nil && opts.Params != nil {
		urlStr = MakeUrl(urlStr, opts.Params)
	}

	// POST
	var reqBody io.Reader
	contentType := ""
	if opts != nil {
		if opts.FormData != nil {
			// 表单数据
			body := &bytes.Buffer{}
			writer := multipart.NewWriter(body)
			for key, val := range opts.FormData {
				_ = writer.WriteField(key, val)
			}
			writer.Close()
			reqBody = body
			contentType = writer.FormDataContentType()
		} else if opts.JSON != nil {
			// JSON请求体
			reqBody = bytes.NewBuffer(opts.JSON)
			contentType = "application/json"
		}
	}

	// 创建 HTTP 请求
	req, err := http.NewRequest(method, urlStr, reqBody)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}

	// 设置 Content-Type 请求头
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}

	// 处理通用选项
	if opts != nil {
		// 设置自定义请求头
		if opts.Headers != nil {
			for key, val := range opts.Headers {
				req.Header.Set(key, val)
			}
		}

		// 设置请求超时
		if opts.Timeout > 0 {
			//ctx, cancel := context.WithTimeout(req.Context(), opts.Timeout)
			//defer cancel()
			//req = req.WithContext(ctx)
			c.client.Timeout = opts.Timeout
		}

		// 设置代理
		if opts.Proxy != "" {
			proxyURL, err := url.Parse(opts.Proxy)
			if err != nil {
				return nil, fmt.Errorf("解析代理URL失败: %w", err)
			}
			c.client.Transport = &http.Transport{Proxy: http.ProxyURL(proxyURL)}
		}
	}

	// 发送 HTTP 请求
	res, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("发送请求失败: %w", err)
	}
	defer res.Body.Close()

	// 读取响应体
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应体失败: %w", err)
	}

	return &Response{res, bodyBytes}, nil
}

// 响应的文本数据
func (r *Response) Text() string {
	return string(r.Body)
}

// 响应的 JSON 数据
func (r *Response) JSON() (map[string]any, error) {
	var jsonMap map[string]any
	err := json.Unmarshal(r.Body, &jsonMap)
	if err != nil {
		return nil, err
	}
	return jsonMap, err
}

// 响应的 JSON 字符串
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

// 响应的 JSON 字符串（适合输出展示）
func (r *Response) PrettyJSONString() (string, error) {
	var buf bytes.Buffer
	err := json.Indent(&buf, r.Body, "", "  ")
	if err != nil {
		return "", fmt.Errorf("格式化JSON失败: %w", err)
	}
	return buf.String(), nil
}
