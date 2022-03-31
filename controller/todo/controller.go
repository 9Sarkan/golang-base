package todo

import (
	"context"
	"time"

	justgrpcservice "github.com/9sarkan/golang-base/api/pb/github.com/9sarkan/golang-base/pkg/pb/justGRPCService.proto"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Controller struct {
	justgrpcservice.UnimplementedToDoListServer
}

func (Controller) Create(ctx context.Context, request *justgrpcservice.CreateToDoRequest) (*justgrpcservice.ToDoSignleResponse, error) {
	return nil, nil
}

func (Controller) Get(ctx context.Context, _ *emptypb.Empty) (*justgrpcservice.ToDoListResponse, error) {
	// mock service for test
	return &justgrpcservice.ToDoListResponse{
		Count: 1,
		Data: []*justgrpcservice.ToDoSignleResponse{
			{
				Id:        uuid.New().String(),
				Title:     "test",
				CreatedAt: time.Now().String(),
				Body:      "test body",
			},
		},
	}, nil
}
