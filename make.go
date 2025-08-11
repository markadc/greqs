package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	_url "net/url"
	"strings"
)

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
