package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("receive request")
		io.WriteString(w, "Hello from Origin Server: 9001")
	})
	log.Println("Origin Server start at :9001")
	log.Fatalln(http.ListenAndServe(":9001", nil))
}
