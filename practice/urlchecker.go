package practice

import (
	"errors"
	"fmt"
	"net/http"
)

type requestResult struct {
	url    string
	status string
}

var errRequestFailed = errors.New("Request failed")

func urlchecker() { // 메인함수는 고루틴을 기다리지 않음. 메인이 먼저 끝나면 남아있는 고루틴도 소멸.
	results := make(map[string]string)
	c := make(chan requestResult)

	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
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
	} // 이야... 확실히 빠르다.

	for url, status := range results {
		fmt.Println(url, status)
	}

}

func hitURL(url string, c chan<- requestResult) { // 이렇게 chan 대신 chan<-을 정해놓으면 보내기만 할 수 있다고 못박아놓을 수 있음.
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		c <- requestResult{url: url, status: "FAILED"}
	} else {
		c <- requestResult{url: url, status: "OK"}
	}
}
