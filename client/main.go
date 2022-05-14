package main

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/jucabet/platzi-protobuffers-grpc/testpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cc, err := grpc.Dial("localhost:5070", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer cc.Close()

	c := testpb.NewTestServiceClient(cc)

	//DoUnary(c)
	//DoClientStreaming(c)
	//DoServerStreaming(c)
	DoBidirectionalStreaming(c)
}

func DoUnary(c testpb.TestServiceClient) {
	req := &testpb.GetTestRequest{
		Id: "t1",
	}

	res, err := c.GetTest(context.Background(), req)
	if err != nil {
		panic(err)
	}

	log.Printf("Response from server: %v", res)
}

func DoClientStreaming(c testpb.TestServiceClient) {
	req := []*testpb.Question{
		{
			Id:       "q8t1",
			Answer:   "Azul",
			Question: "color asociado a golang?",
			TestId:   "t1",
		},
		{
			Id:       "q9t1",
			Answer:   "google",
			Question: "Empresa que desarrollo golang",
			TestId:   "t1",
		},
		{
			Id:       "q10t1",
			Answer:   "Backend",
			Question: "Especialidad de golang",
			TestId:   "t1",
		},
	}

	stream, err := c.SetQuestions(context.Background())
	if err != nil {
		panic(err)
	}

	for _, question := range req {
		log.Println("Sending question: ", question.Id)
		stream.Send(question)
		time.Sleep(2 * time.Second)
	}

	msg, err := stream.CloseAndRecv()
	if err != nil {
		panic(err)
	}

	log.Printf("Response from server: %v", msg)
}

func DoServerStreaming(c testpb.TestServiceClient) {
	req := &testpb.GetStudentsPerTestRequest{
		TestId: "t1",
	}

	stream, err := c.GetStudentsPerTest(context.Background(), req)
	if err != nil {
		panic(err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}
		log.Printf("Response From server: %v", msg)
	}
}

func DoBidirectionalStreaming(c testpb.TestServiceClient) {
	answer := testpb.TakeTestRequests{
		Answer: "42",
	}

	nuberOfQuestions := 4

	waitChannel := make(chan struct{})

	stream, err := c.TakeTest(context.Background())
	if err != nil {
		panic(err)
	}

	go func() {
		for i := 0; i < nuberOfQuestions; i++ {
			stream.Send(&answer)
			time.Sleep(2 * time.Second)
		}
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				panic(err)
			}

			log.Printf("Response From server: %v", res)
		}
		close(waitChannel)
	}()

	<-waitChannel
}
