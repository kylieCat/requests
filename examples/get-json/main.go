package main

import (
	"fmt"
	"log"

	"github.com/kylie-a/requests"
	"bytes"
	"encoding/json"
)

func main() {
	var client requests.Client
	resp, err := client.Get("https://httpbin.org/get")
	check(err)

	m := make(map[string]interface{})
	err = resp.JSON(&m)
	check(err)

	fmt.Printf("%#v\n", m)

	j , _ := json.Marshal(m)
	resp, err = client.Post("https://jsonplaceholder.typicode.com/posts/", bytes.NewReader(j))
	check(err)
	err = resp.JSON(&m)
	check(err)

	fmt.Println(m)
}

func check(err error) {
	if err != nil {
		log.Fatalf("%+v", err)
	}
}
