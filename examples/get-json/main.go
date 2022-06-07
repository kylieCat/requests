package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"github.com/kyliecat/requests"
)

func main() {
	client := requests.NewClient()
	resp, err := client.Get("https://httpbin.org/get")
	check(err)

	m := make(map[string]interface{})
	err = resp.JSON(&m)
	check(err)

	fmt.Printf("%#v\n", m)

	j, _ := json.Marshal(m)
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
