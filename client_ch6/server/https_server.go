package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"time"
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

func handlerUpgrade(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Connection") != "Upgrade" || r.Header.Get("Upgrade") != "MyProtocol" {
		w.WriteHeader(400)
		return
	}
	fmt.Println("Upgrade to MyProtocol")

	hijacker := w.(http.Hijacker)
	conn, readWriter, err := hijacker.Hijack()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	res := http.Response{
		StatusCode: 101,
		Header:     make(http.Header),
	}
	res.Header.Set("Upgrade", "MyProtocol")
	res.Header.Set("Connection", "Upgrade")
	res.Write(conn)

	for i := 1; i <= 10; i++ {
		fmt.Fprintf(readWriter, "%d\n", i)
		fmt.Println("->", i)
		readWriter.Flush()
		recv, err := readWriter.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		fmt.Printf("<- %s", string(recv))
		time.Sleep(500 * time.Millisecond)
	}
}

//test update
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
