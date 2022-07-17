package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"go.mondoo.com/ranger-rpc/examples/pingpong"
)

func main() {
	fmt.Println("start pingpong example client")
	client, err := pingpong.NewPingPongClient("http://localhost:2155/api/", &http.Client{
		Timeout: time.Second * 10,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("call server ping")
	resp, err := client.Ping(context.Background(), &pingpong.PingRequest{Sender: "bob"})
	if err == nil {
		fmt.Println("server ping response:", resp.Message)
	} else {
		panic(err)
	}
}
