package main

import (
	"context"
	"net/http"

	"godaily/proto"
)

type Server struct{}

func (s *Server) Say(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	return &proto.Response{Text: request.GetText()}, nil
}

func main() {
	server := &Server{}
	twirpHandler := proto.NewEchoServer(server, nil)

	http.ListenAndServe(":8080", twirpHandler)
}
