package greqs

import (
	"net/http"
	"time"
)

type Worker struct {
	proxy       string
	timeout     time.Duration
	requestHook func(req *http.Request)
	proxyHook   func(cli *http.Client)
}

func NewWorker(proxy string, timeout time.Duration, reqHook func(req *http.Request), proxyHook func(cli *http.Client)) *Worker {
	return &Worker{
		proxy:       proxy,
		timeout:     timeout,
		requestHook: reqHook,
		proxyHook:   proxyHook,
	}
}

func (w *Worker) GetProxy() string {
	return w.proxy
}

func (w *Worker) SetProxy(proxy string) {
	w.proxy = proxy
}

func (w *Worker) GetTimeout() time.Duration {
	return w.timeout
}

func (w *Worker) SetTimeout(timeout time.Duration) {
	w.timeout = timeout
}

func (w *Worker) Get(url string, headers S) (*Response, error) {
	req, err := MakeGetRequest(url, headers)
	if err != nil {
		return nil, err
	}
	return w.Go(req)
}

func (w *Worker) Post(url string, headers S, data A) (*Response, error) {
	req, err := MakePostRequest(url, headers, data)
	if err != nil {
		return nil, err
	}
	return w.Go(req)
}

func (w *Worker) PostForm(url string, headers S, form S) (*Response, error) {
	req, err := MakePostFormRequest(url, headers, form)
	if err != nil {
		return nil, err
	}
	return w.Go(req)
}

func (w *Worker) Go(req *http.Request) (*Response, error) {
	if w.requestHook != nil {
		w.requestHook(req)
	}

	cli := GetClient(w.proxy, w.timeout)

	if w.proxyHook != nil {
		w.proxyHook(cli)
	}
	return Do(cli, req)
}
