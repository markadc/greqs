package greqs

import (
	"fmt"
	"greqs/log"
	"testing"
)

var (
	//ua      = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/138.0.0.0 Safari/537.36 Edg/138.0.0.0"
	ua      = "Greqs"
	headers = S{"User-Agent": ua}
)

func TestMakeUrl(t *testing.T) {
	params := S{"name": "Greqs", "type": "GoLang"}

	url1 := "https://httpbin.org/get"
	fmt.Println(MakeUrl(url1, params))

	url2 := "https://httpbin.org/get?"
	fmt.Println(MakeUrl(url2, params))

}

func TestGet(t *testing.T) {
	url := "https://httpbin.org/get"
	resp, _ := Get(url, headers)

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Text())
}

func TestPost(t *testing.T) {
	url := "https://httpbin.org/post"
	data := A{"name": "Greqs", "type": "GoLang"}
	resp, _ := Post(url, data, headers)

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Text())
}

func TestPostForm(t *testing.T) {
	url := "https://httpbin.org/post"
	form := S{"name": "Greqs", "type": "GoLang"}
	resp, _ := PostForm(url, form, nil)

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Text())
}

func TestResponse(t *testing.T) {
	url := "https://httpbin.org/get"
	resp, _ := Get(url, headers)

	log.Red("Call...JSON")
	j, _ := resp.JSON()
	fmt.Printf("%s\n\n", j)

	log.Red("Call...JSONString")
	js, _ := resp.JSONString()
	fmt.Printf("%s\n\n", js)

	log.Red("Call...PrettyJSONString")
	pjs, _ := resp.PrettyJSONString()
	fmt.Printf("%s\n\n", pjs)
}

func TestSend(t *testing.T) {
	url := "https://httpbin.org/get"
	res, _ := Send("GET", url, &Options{Headers: headers})
	fmt.Println(res.StatusCode)

}
