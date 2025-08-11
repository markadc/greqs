# 项目说明

一个 golang 的网络请求库

# 使用

## 发送 GET 请求

```go
import (
    "fmt"
    "greqs/requests"
)

func main() {
worker := requests.NewWorker()

params := map[string]string{"page": "1", "limit": "10"}
headers := map[string]string{"User-Agent": "Greqs"}
opts := &requests.Options{Params: params, Headers: headers}

resp, _ := worker.Get("https://httpbin.org/get", opts)

fmt.Println(resp.StatusCode)
fmt.Println(resp.Text())
}
```

## 发送 POST 请求

### 表单数据

```golang
import (
    "fmt"
    "greqs/requests"
)

func main() {
worker := requests.NewWorker()

formData := map[string]string{"name": "form"}
opts := &requests.Options{FormData: formData}

resp, _ := worker.Post("https://httpbin.org/post", opts)

fmt.Println(resp.StatusCode)
fmt.Println(resp.Text())
}

```

### JSON 请求体

```go
import (
    "encoding/json"
    "fmt"
    "greqs/requests"
)

func main() {
worker := requests.NewWorker()

jsonMap := map[string]string{"name": "json"}
b, _ := json.Marshal(jsonMap)
opts := &requests.Options{JSON: b}

resp, _ := worker.Post("https://httpbin.org/post", opts)

fmt.Println(resp.StatusCode)
fmt.Println(resp.Text())
}

```

## 使用代理

```go


import (
    "fmt"
    "greqs/requests"
)

func main() {
worker := requests.NewWorker()

opts := &requests.Options{Proxy: "http://127.0.0.1:7890"}

resp, err := worker.Get("https://httpbin.org/ip", opts)
if err != nil {
fmt.Printf("Error => %s\n", err)
return
}

fmt.Println(resp.StatusCode)
fmt.Println(resp.Text())
}

```

## 设置超时

```python

```
