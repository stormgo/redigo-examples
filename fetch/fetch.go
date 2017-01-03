package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	url := "http://127.0.0.1:3000/omdb.json"
	json := getJson(url)
	fmt.Println(len(json))
	reader := bytes.NewReader(json)

	scanner := bufio.NewScanner(reader)
	count := 0
	for scanner.Scan() {
		evenodd := count % 2
		if evenodd == 0 {
			fmt.Println(count, " ", scanner.Text()) // Println will add back the final '\n'
		} else {
			fmt.Println(count, " ", scanner.Text()) // Println will add back the final '\n'
		}
		count = count + 1
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
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
