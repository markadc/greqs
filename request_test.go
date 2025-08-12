package greqs

import (
	"fmt"
	"testing"
	"time"
)

func TestTask_Do(t *testing.T) {
	req := &Request{
		Method:  "GET",
		Url:     "https://httpbin.org/get",
		Headers: map[string]string{"Name": "Greqs"},
		Timeout: 3 * time.Second,
	}
	resp, err := req.Do()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(resp.Text())
}
