package main

import (
	"net"

	"github.com/jucabet/platzi-protobuffers-grpc/tree/main/database"
	"github.com/jucabet/platzi-protobuffers-grpc/tree/main/server"
	"github.com/jucabet/platzi-protobuffers-grpc/tree/main/studentpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	list, err := net.Listen("tcp", ":5060")
	if err != nil {
		panic(err)
	}

	repo, err := database.NewPostgresRepository("postgres://postgres:postgres@localhost:54321/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}

	server := server.NewStudentServer(repo)

	s := grpc.NewServer()
	studentpb.RegisterStudentServiceServer(s, server)

	reflection.Register(s)

	if err := s.Serve(list); err != nil {
		panic(err)
	}
}
