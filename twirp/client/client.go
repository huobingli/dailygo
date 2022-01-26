package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"godaily/proto"
)

func main() {
	client := proto.NewEchoProtobufClient("http://localhost:8080", &http.Client{})

	response, err := client.Say(context.Background(), &proto.Request{Text: "Hello world"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("response:%s\n", response.GetText())
}
