package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func handler(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dump))
	fmt.Fprintf(w, "<html><body>hello</body></html>\n")
}

func highLevelApi() {
	http.HandleFunc("/", handler)
	log.Println("start http listening: 18443")
	err := http.ListenAndServeTLS(":18443", "./cert/server.crt", "./cert/server.key", nil)
	log.Println(err)
}

func http_Server() {
	server := &http.Server{
		TLSConfig: &tls.Config{
			ClientAuth: tls.RequireAndVerifyClientCert,
			MinVersion: tls.VersionTLS12,
		},
		Addr: ":18443",
	}
	http.HandleFunc("/", handler)
	log.Println("start http listening: 18443")
	err := server.ListenAndServeTLS("./cert/server.crt", "./cert/server.key")
	log.Println(err)
}

func main() {
	//highLevelApi()
	http_Server()
}
