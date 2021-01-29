package main

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"time"
)

var image []byte
var html []byte

// send HTML to browser
func handlerHtml(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	w.Write(html)
}

func handlerPrimeSSE(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}
	ctx := r.Context()
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var num int64 = 1
	for id := 1; id <= 100; id++ {
		select {
		case <-ctx.Done():
			fmt.Println("connection closed from client")
			return
		default:
			//do noting
		}
		for {
			num++
			if big.NewInt(num).ProbablyPrime(20) {
				fmt.Println(num)
				fmt.Fprintf(w, "data: {\"id\": %d, \"number\": %d}\n\n", id, num)
				flusher.Flush()
				time.Sleep(time.Second)
				break
			}
		}
		time.Sleep(time.Second)
	}
	fmt.Println("Connection colsed from server")
}

func init() {
	var err error
	image, err = ioutil.ReadFile("./image.jpeg")
	if err != nil {
		panic(err)
	}
}

// func handlerHtml(w http.ResponseWriter, r *http.Request) {
// 	pusher, ok := w.(http.Pusher)
// 	if ok {
// 		pusher.Push("/image", nil)
// 	}
// 	w.Header().Add("Content-Type", "text/html")
// 	fmt.Fprintf(w, `<html><body><img src="/image"></body></html>`)
// }

func handlerImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/jpeg")
	w.Write(image)
}

func main() {
	// http.HandleFunc("/", handlerHtml)
	// http.HandleFunc("/image", handlerImage)
	// fmt.Println("start http listening :18443")
	// err := http.ListenAndServeTLS(":18443", "./cert/server.crt", "./cert/server.key", nil)
	// fmt.Println(err)
	var err error
	html, err = ioutil.ReadFile("index.html")
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/", handlerHtml)
	http.HandleFunc("/prime", handlerPrimeSSE)
	fmt.Println("start http listening: 18888")
	err = http.ListenAndServe(":18888", nil)
	fmt.Println(err)
}
