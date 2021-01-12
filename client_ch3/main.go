package main

import (
	"log"
	"net/http"
	"net/url"
)

func post() {
	values := url.Values{
		"test": {"values"},
	}
	resp, err := http.PostForm("http://localhost:18888", values)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	log.Println("Status:", resp.Status)

}

func main() {

	post()
}
