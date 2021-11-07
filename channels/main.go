package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://google.com",
		"http://amazon.com",
		"http://uol.com.br",
		"http://prefeitura.sp.gov.br",
		"http://aprogeosp.org.br",
		"http://zap.com.br",
		"http://vivareal.com.br",
	}
	channel := make(chan string)

	for _, link := range links {
		go checkLink(link, channel)
	}

	for link := range channel {
		go func(l string) {
			time.Sleep(time.Second * 30)
			checkLink(l, channel)
		}(link)
	}
}

func checkLink(link string, channel chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "Might be down!")
		channel <- link
		return
	}
	fmt.Println(link, "yes, it's up!")
	channel <- link
}
