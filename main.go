package main

import (
	"fmt"
	"net/http"
)

type requestResult struct {
	url    string
	status string
}

func main() {
	var results = make(map[string]string)
	c := make(chan requestResult) // result를 보내고 받는 채널 생성

	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}

	for _, url := range urls {
		go hitURL(url, c)
	}

	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}
}

func hitURL(url string, c chan requestResult) {
	resp, err := http.Get(url)
	status := "OK"

	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- requestResult{url: url, status: status}
}
