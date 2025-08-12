package greqs

// Get 发送 GET 请求
func Get(url string, headers S) (*Response, error) {
	req := &Request{
		Method:  "GET",
		Url:     url,
		Headers: headers,
	}
	return req.Do()
}

// POST 发送 POST 请求（JOSN 形式）
func Post(url string, headers S, data A) (*Response, error) {
	req := &Request{
		Method:  "POST",
		Url:     url,
		Headers: headers,
		Data:    data,
	}
	return req.Do()
}

// PostForm 发送 POST 请求（Form 形式）
func PostForm(url string, headers S, form S) (*Response, error) {
	req := &Request{
		Method:  "POST",
		Url:     url,
		Headers: headers,
		Form:    form,
	}
	return req.Do()
}
