package main

import (
	"fmt"
	"log"

	"github.com/kyliecat/requests"
)

func main() {
	client := requests.NewClient()
	resp, err := client.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatalf("%+v", err)
	}
	fmt.Println(resp.Request.Method, resp.Request.URL, resp.Status.Code)
}
