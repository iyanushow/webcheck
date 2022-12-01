package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://amazon.com",
		"https://golang.org",
	}

	c := make(chan string)

	for _, link := range links {
		go checkWebsite(link, c)
	}

	for l := range c {
		go func(lk string) {
			time.Sleep(time.Second)
			checkWebsite(lk, c)
		}(l)
	}
}

func checkWebsite(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down")
		c <- link
		return
	}

	fmt.Println(link, "is up ")
	c <- link
}
