package main

import (
	"context"
	"log"
	"net"

	"sample-manager/config"
	"sample-manager/model"
	pb "sample-manager/proto"

	"github.com/lib/pq"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedSampleServiceServer
}

func (s *server) GetSampleItemID(ctx context.Context, in *pb.GetSampleItemIDRequest) (*pb.GetSampleItemIDResponse, error) {
	var sampleItem model.SampleItem
	if err := config.GetDB().Where("item_id = ? AND segments = ?::text[]", in.ItemId, pq.Array(in.ClmSegments)).First(&sampleItem).Error; err != nil {
		return nil, err
	}

	sampleItemID := sampleItem.SampleItemID

	return &pb.GetSampleItemIDResponse{SampleItemId: sampleItemID}, nil
}

func (s *server) CreateSampleItem(ctx context.Context, in *pb.CreateSampleItemRequest) (*pb.CreateSampleItemResponse, error) {
	tx := config.GetDB().Begin()

	sampleItem := &model.SampleItem{
		SampleItemID: in.SampleItemId,
		ItemID:       in.ItemId,
		Segments:     in.ClmSegments,
	}

	if err := tx.Create(sampleItem).Error; err != nil {
		tx.Rollback()
		return nil, err
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
