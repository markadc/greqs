package greqs

import (
	"fmt"
	"testing"
)

func TestTask_Do(t *testing.T) {
	req := &Request{
		Method:  "GET",
		Url:     "https://httpbin.org/get",
		Headers: map[string]string{"Name": "Greqs"},
	}
	resp, _ := req.Do()
	fmt.Println(resp.Text())
}
