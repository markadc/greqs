package greqs

import (
	"fmt"
	"strings"
	"time"
)

// Request 请求
type Request struct {
	Method  string
	Url     string
	Params  S
	Headers S
	Data    A
	Form    S
	Proxy   string
	Timeout time.Duration
}

// Do 执行请求
func (r *Request) Do() (*Response, error) {
	method := strings.ToUpper(r.Method)
	if method != "GET" && method != "POST" {
		return nil, fmt.Errorf("仅支持 GET、POST，不支持 %s", method)
	}

	if r.Params != nil {
		r.Url = MakeUrl(r.Url, r.Params)
	}

	cli := GetClient(r.Proxy, r.Timeout)

	if method == "GET" {
		req, err := MakeGetRequest(r.Url, r.Headers)
		if err != nil {
			return nil, err
		}
		return Do(cli, req)

	} else {
		if r.Data != nil {
			req, err := MakePostRequest(r.Url, r.Headers, r.Data)
			if err != nil {
				return nil, err
			}
			return Do(cli, req)

		} else if r.Form != nil {
			req, err := MakePostFormRequest(r.Url, r.Headers, r.Form)
			if err != nil {
				return nil, err
			}
			return Do(cli, req)
		} else {
			return nil, fmt.Errorf("无效的 POST 请求")
		}
	}
}
