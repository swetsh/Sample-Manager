package main

import (
	"context"
	"log"
	"net"

	"sample-manager/config"
	pb "sample-manager/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedSampleManagerServer
}

func (s *server) GetSampleItemID(ctx context.Context, in *pb.GetSampleItemIDRequest) (*pb.GetSampleItemIDResponse, error) {
	return &pb.GetSampleItemIDResponse{SampleItemId: "sample_item_id_123"}, nil
}

func main() {
	config.DatabaseConnection()

	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSampleManagerServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
