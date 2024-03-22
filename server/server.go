package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"sample-manager/config"
	"sample-manager/model"
	pb "sample-manager/proto"

	"github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type Server struct {
	pb.SampleServiceServer
	DB *gorm.DB
}

func (s *Server) GetSampleItemID(ctx context.Context, in *pb.GetSampleItemIDRequest) (*pb.GetSampleItemIDResponse, error) {
	var sampleItem model.SampleItem
	if err := s.DB.Where("item_id = ? AND segments = ?::text[]", in.ItemId, pq.Array(in.ClmSegments)).First(&sampleItem).Error; err != nil {
		errorString := fmt.Sprintf("No mapping found: %v", err)
		return nil, status.Error(codes.Unavailable, errorString)
	}

	sampleItemID := sampleItem.SampleItemID

	return &pb.GetSampleItemIDResponse{SampleItemId: sampleItemID}, nil
}

func (s *Server) CreateSampleItem(ctx context.Context, in *pb.CreateSampleItemRequest) (*pb.CreateSampleItemResponse, error) {
	tx := s.DB.Begin()

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
	pb.RegisterSampleServiceServer(s, &Server{DB: config.GetDB()})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
