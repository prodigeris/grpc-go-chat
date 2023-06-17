package main

import (
	"context"
	"github.com/prodigeris/grpc-go-chat/pb"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
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
	arnas := &pb.Participant{Id: "1", Name: "Arnas"}
	petras := &pb.Participant{Id: "2", Name: "Petras"}
	return []*pb.Message{
		&pb.Message{
			Text:   "Labas, as Arnas",
			Author: arnas,
			Date:   &timestamppb.Timestamp{Seconds: 1687023330},
		},
		&pb.Message{
			Text:   "How are you?",
			Author: arnas,
			Date:   &timestamppb.Timestamp{Seconds: 1687023340},
		},
		&pb.Message{
			Text:   "I'm good!",
			Author: petras,
			Date:   &timestamppb.Timestamp{Seconds: 1687023342},
		},
	}
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
