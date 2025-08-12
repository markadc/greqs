package greqs

import (
	"fmt"
	"testing"
)

func TestReqs(t *testing.T) {
	// GET
	url := "https://httpbin.org/get"
	r1, _ := Get(url, nil)
	fmt.Println(r1.Text())

	// POST JOSN
	url = "https://httpbin.org/post"
	data := A{"Type": "JSON", "Value": []int{1, 2, 3}}
	r2, _ := Post(url, nil, data)
	fmt.Println(r2.Text())

	// POST From
	url = "https://httpbin.org/post"
	form := S{"Type": "From"}
	r3, _ := PostForm(url, nil, form)
	fmt.Println(r3.Text())
}
