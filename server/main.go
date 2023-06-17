package main

import (
	"context"
	"github.com/prodigeris/grpc-go-chat/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedTranscriptServer
}

func (s *server) GetBookList(ctx context.Context, in *pb.GetMessageListRequest) (*pb.GetMessageListResponse, error) {
	return &pb.GetMessageListResponse{
		Messages: getSampleMessages(),
	}, nil
}

func getSampleMessages() []*pb.Message {
	return []*pb.Message{}
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterTranscriptServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
