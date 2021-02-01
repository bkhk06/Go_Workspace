// poll_url project main.go
package main

import (
	"fmt"
	"net/http"
)

var urls = []string{
	"http://www.baidu.com/",
	"http://www.sina.com/",
	"http://blog.sina.com/",
}

func main() {
	// Execute an HTTP HEAD request for all url's
	// and returns the HTTP status string or an error string.
	for _, url := range urls {
		resp, err := http.Head(url)
		if err != nil {
			fmt.Println("Error:", url, err)
		}
		fmt.Println(url, ": ", resp.Status)
	}
}
