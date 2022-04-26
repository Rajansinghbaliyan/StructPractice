package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func practiceInterface() {
	resp, err := http.Get("https://google.com")
	//content := make([]byte, 99999)
	//_, err = resp.Body.Read(content)

	//fmt.Println(string(content))

	//err = ioutil.WriteFile("google.html", content, 0666)

	writeData := writToFile{}
	//io.Copy(os.Stdout, resp.Body)
	_, err = io.Copy(writeData, resp.Body)

	if err != nil {
		fmt.Printf("%v", err)
	}

	err = resp.Body.Close()

	//fmt.Printf("%+v", string(writeData))
}

type writToFile struct {
}

func (writToFile) Write(receivingData []byte) (int, error) {
	//f.data = append(f.data, receivingData...)
	//fmt.Println(string(receivingData))
	err := ioutil.WriteFile("google.html", receivingData, 0666)
	return len(receivingData), err
}
