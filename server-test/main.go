package main

import (
	"net"

	"github.com/jucabet/platzi-protobuffers-grpc/tree/main/database"
	"github.com/jucabet/platzi-protobuffers-grpc/tree/main/server"
	"github.com/jucabet/platzi-protobuffers-grpc/tree/main/testpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	list, err := net.Listen("tcp", ":5070")
	if err != nil {
		panic(err)
	}

	repo, err := database.NewPostgresRepository("postgres://postgres:postgres@localhost:54321/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}

	server := server.NewTestServer(repo)

	s := grpc.NewServer()
	testpb.RegisterTestServiceServer(s, server)

	reflection.Register(s)

	if err := s.Serve(list); err != nil {
		panic(err)
	}
}
