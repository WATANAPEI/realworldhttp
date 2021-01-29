package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	client := &http.Client{}
	resp, err := client.Get("https://api.github.com/user/emails")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	emails, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(emails))

}
