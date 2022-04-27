package main

import (
	"fmt"
	"net/http"
)

var urls = []string{
	"https://google.com",
	"https://amazon.com",
	"https://yahoo.com",
	"https://youtube.com",
	"https://facebook.com",
	"https://flipkart.com",
}

func getContentFromUrl(url string) (result string, err error) {
	//fmt.Printf("Inside the getContentFromUrl: %v\n", url)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	if resp.StatusCode == 200 {
		result = fmt.Sprintf("This site %v is up", url)
	} else {
		result = fmt.Sprintf("This site %v is down", url)
	}
	err = resp.Body.Close()
	return
}

func fetchingContent() {
	for _, url := range urls {
		result, err := getContentFromUrl(url)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(result)
		}
	}
}

func addToChannel(
	resultMapCh chan string,
	url string,
	resultFetched *int,
) {
	//fmt.Printf("Inside the addTChannel: %v\n", url)
	result, err := getContentFromUrl(url)
	if err != nil {
		*resultFetched = *resultFetched + 1
		resultMapCh <- fmt.Sprint(err)
	} else {
		*resultFetched = *resultFetched + 1
		resultMapCh <- result
	}
}

func fetchingContentConcurrently() {
	resultCh := make(chan string)
	resultFetched := 0

	resultMap := make(map[int]string)

	for _, url := range urls {
		go addToChannel(resultCh, url, &resultFetched)
	}

	//for resultFetched < len(urls) {
	//	fmt.Println("For loop")
	//	time.Sleep(3000)
	//}

	for i := range urls {
		resultMap[i] = <-resultCh
	}

	// can't have extra listener than writer
	//fmt.Println(<-resultCh)

	for _, value := range resultMap {
		fmt.Println(value)
	}

}
