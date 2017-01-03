package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	url := "http://127.0.0.1:3000/omdb.json"
	json := getJson(url)
	fmt.Println(len(json))
}

func getJson(url string) (buf []byte) {
	var netClient = &http.Client{
		Timeout: time.Second * 30,
	}
	response, err := netClient.Get(url)
	if err != nil {
		fmt.Println("get: ", err)
	}

	buf, err = ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("ioutil: ", err)
	}
	return buf
}
