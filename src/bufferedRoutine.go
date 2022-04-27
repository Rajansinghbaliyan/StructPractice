package main

import (
	"fmt"
	"net/http"
	"time"
)

//var urls = []string{
//	"https://google.com",
//	"https://amazon.com",
//	"https://yahoo.com",
//	"https://youtube.com",
//	"https://facebook.com",
//	"https://flipkart.com",
//}

func main() {
	//unbuffered()
	//buffered()
	//channelSelection()
	channelSelectionWithTimeout()
}

func unbuffered() {
	channel := make(chan int)

	// Will create a deadlock because there will be no listener
	//channel <- 2
	// These children will be blocked if there is no listener
	go func() { channel <- 1 }()
	go func() { channel <- 2 }()
	go func() { channel <- 3 }()

	first := <-channel
	second := <-channel
	third := <-channel

	fmt.Println(first, second, third)
}

func buffered() {
	channel := make(chan int, 2)

	channel <- 1
	channel <- 2

	// will also create a deadlock because the buffered capacity is 2
	// channel <- 3

	go func() { channel <- 4 }()

	first := <-channel
	second := <-channel
	third := <-channel

	fmt.Println(first, second, third)

}

func channelSelection() {
	respCh := make(chan string)
	errorCh := make(chan string)

	for _, url := range urls {
		go getDataFromUrl(respCh, errorCh, url)
	}

	for {
		select {
		case resp := <-respCh:
			fmt.Println(resp)
		case err := <-errorCh:
			fmt.Println(err)
		default:
			fmt.Println("No data fetched yet.")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func getDataFromUrl(respCh chan string, errorCh chan string, url string) {
	_, err := http.Get(url)
	if err != nil {
		errorCh <- fmt.Sprint(err)
	} else {
		respCh <- fmt.Sprintf("%v is active.", url)
	}
}

func channelSelectionWithTimeout() {
	respCh := make(chan string)
	errorCh := make(chan string)

	for _, url := range urls {
		go getDataFromUrl(respCh, errorCh, url)
	}

	for {
		select {
		case resp := <-respCh:
			fmt.Println(resp)
		case err := <-errorCh:
			fmt.Println(err)
		case <-time.After(3000 * time.Millisecond):
			fmt.Println("Time out")
			return
		}
	}
}
