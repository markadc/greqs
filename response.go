package greqs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

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
