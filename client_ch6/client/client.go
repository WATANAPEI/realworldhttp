package main

import (
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	res, err := http.Get("https://localhost:18443")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	dump, err := httputil.DumpResponse(res, true)
	if err != nil {
		panic(err)
	}
	log.Println(string(dump))
}
