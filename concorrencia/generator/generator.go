package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// <-chan - canal somente-leitura
func titulo(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			c <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}

	return c
}

func main() {
	t1 := titulo("https://www.udemy.com", "https://calendar.google.com")
	t2 := titulo("https://mail.google.com", "https://youtube.com")
	fmt.Println(<-t1, <-t2)
	fmt.Println(<-t1, <-t2)
}
