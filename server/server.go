package main

import (
	"context"
	"log"
	"net"

	"sample-manager/config"
	"sample-manager/model"
	pb "sample-manager/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedSampleServiceServer
}

func (s *server) GetSampleItemID(ctx context.Context, in *pb.GetSampleItemIDRequest) (*pb.GetSampleItemIDResponse, error) {
	return &pb.GetSampleItemIDResponse{SampleItemId: "sample_item_id_123"}, nil
}

func (s *server) CreateSampleItem(ctx context.Context, in *pb.CreateSampleItemRequest) (*pb.CreateSampleItemResponse, error) {
	tx := config.GetDB().Begin()

	sampleItem := &model.SampleItem{
		SampleItemID: in.SampleItemId,
		ItemID:       in.ItemId,
	}

	if err := tx.Create(sampleItem).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	for _, segment := range in.ClmSegments {
		newSegment := &model.Segment{
			Segment:   segment,
			MappingID: sampleItem.ID,
		}
		if err := tx.Create(newSegment).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return &pb.CreateSampleItemResponse{Message: "Sample item created successfully"}, nil
}

func main() {
	config.DatabaseConnection()

	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSampleServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
