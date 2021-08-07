package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://youtube.com",
		"http://golang.org",
	}

	c := make(chan string)

	for _, link := range links {
		//go keyword creates new goroutines specifically for the checkLink function
		go checkLink(link, c) //this works one by one //very slow if we have many websites
	}
	//for i := 0; i<len(links); i++
	//for //infinite loop
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l) //value of channel printing is a blocking call
	}
}

//go keyword is used only in front of function calls
//if main routine ends before any child routine ends, the whole program ends
//channels can be used to communicate between 2 goroutines : main and child

//a go scheduler is linked with one core of our CPU by default, we can change it to multiple cores as per our will
//concurrency - (one core)multiple threads executing code. ig one thread blocked, another one is picked up and worked on
//parallelism - (multiple cores)multiple threads executed at the exact same time
//picking up one and getting back to it vs running some at same time

func checkLink(link string, c chan string) {
	_, err := http.Get(link) //get call is a blocking call and main goroutine is blocked
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
