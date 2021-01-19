package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"net"
	"net/http"
	"strconv"
	"time"
)

func easyChunk() {
	resp, err := http.Get("http://localhost:18888/chunked")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	reader := bufio.NewReader(resp.Body)
	for {
		line, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		log.Println(string(bytes.TrimSpace(line)))
	}
}

func completeChunk() {
	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}
	conn, err := dialer.Dial("tcp", "localhost:18888")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	request, err := http.NewRequest("GET", "http://localhost:18888/chunked", nil)
	err = request.Write(conn)
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(conn)
	resp, err := http.ReadResponse(reader, request)
	if err != nil {
		panic(err)
	}
	if resp.TransferEncoding[0] != "chunked" {
		panic("wrong transfer encoding")
	}
	for {
		sizeStr, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		size, err := strconv.ParseInt(string(sizeStr[:len(sizeStr)-2]), 16, 64)
		if size == 0 {
			break
		}
		if err != nil {
			panic(err)
		}
		line := make([]byte, int(size))
		reader.Read(line)
		reader.Discard(2)
		log.Println(" ", string(line))
	}

}

func main() {
	completeChunk()

}
