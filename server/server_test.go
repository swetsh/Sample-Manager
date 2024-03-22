package main

import (
	"context"
	pb "sample-manager/proto"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestCreatingSampleItem(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.Nil(t, err, "Error creating mock db: %v", err)

	defer db.Close()

	dialect := postgres.New(postgres.Config{
		Conn:       db,
		DriverName: "postgres",
	})

	gormDb, err := gorm.Open(dialect, &gorm.Config{})
	assert.Nil(t, err, "Error creating mock gorm db: %v", err)

	type args struct {
		ctx context.Context
		req *pb.CreateSampleItemRequest
	}

	tests := []struct {
		name      string
		args      args
		rows      func()
		want      *pb.CreateSampleItemResponse
		wantErr   bool
		errorCode codes.Code
	}{
		{
			name: "Creating a sampleitem - Success",
			args: args{
				ctx: context.Background(),
				req: &pb.CreateSampleItemRequest{

					ClmSegments:  []string{"seg1", "seg2"},
					ItemId:       "itemid",
					SampleItemId: "sampleid",
				},
			},
			rows: func() {
				mock.ExpectBegin()
				rows := sqlmock.NewRows([]string{
					"id",
					"segments",
					"item_id",
					"sample_item_id",
				}).AddRow(1, pq.StringArray{"seg1", "seg2"}, "itemid", "sampleid")
				mock.ExpectQuery("INSERT").WillReturnRows(rows)
				mock.ExpectCommit()
			},
			want: &pb.CreateSampleItemResponse{
				Message: "Sample item created successfully",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.rows()
			server := &Server{DB: gormDb}

			got, err := server.CreateSampleItem(tt.args.ctx, tt.args.req)

			if (err != nil) != tt.wantErr {
				t.Fatalf("CreateMapping() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				statusErr, ok := status.FromError(err)
				assert.True(t, ok, "Expected gRPC status error")
				assert.Equalf(t, tt.errorCode, statusErr.Code(), "Expected %v error", tt.errorCode)
			} else {
				assert.Equalf(t, tt.want, got, "CreateMapping(%v, %v)", tt.args.ctx, tt.args.req)
			}
		})
	}
}

func TestGettingASampleID(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.Nil(t, err, "Error creating mock db: %v", err)

	defer db.Close()

	dialect := postgres.New(postgres.Config{
		Conn:       db,
		DriverName: "postgres",
	})

	gormDb, err := gorm.Open(dialect, &gorm.Config{})
	assert.Nil(t, err, "Error creating mock gorm db: %v", err)

	type args struct {
		ctx context.Context
		req *pb.GetSampleItemIDRequest
	}
	tests := []struct {
		name      string
		args      args
		rows      func()
		want      *pb.GetSampleItemIDResponse
		wantErr   bool
		errorCode codes.Code
	}{
		{
			name: "Getting a sample ID - Expect Success",
			args: args{
				ctx: context.Background(),
				req: &pb.GetSampleItemIDRequest{
					ClmSegments: []string{"seg1", "seg2"},
					ItemId:      "itemid",
				},
			},
			rows: func() {
				mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{"id", "sample_item_id", "item_id", "segments"}).
					AddRow(1, "1", "itemid", pq.StringArray{"seg1", "seg2"}))
			},
			want: &pb.GetSampleItemIDResponse{
				SampleItemId: "1",
			},
			wantErr: false,
		},
		{
			name: "Getting a sample ID when mapping doesn't exist - Expect Error",
			args: args{
				ctx: context.Background(),
				req: &pb.GetSampleItemIDRequest{
					ClmSegments: []string{"seg1", "seg2"},
					ItemId:      "itemid",
				},
			},
			rows: func() {
				mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{}))
			},
			want:      nil,
			wantErr:   true,
			errorCode: codes.Unavailable,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.rows()
			server := &Server{DB: gormDb}

			got, err := server.GetSampleItemID(tt.args.ctx, tt.args.req)

			if (err != nil) != tt.wantErr {
				t.Fatalf("GetSampleId() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				statusErr, ok := status.FromError(err)
				assert.True(t, ok, "Expected gRPC status error")
				assert.Equalf(t, tt.errorCode, statusErr.Code(), "Expected %v error", tt.errorCode)
			} else {
				assert.Equalf(t, tt.want, got, "GetSampleId(%v, %v)", tt.args.ctx, tt.args.req)
			}
		})
	}
}
