package todo

import (
	"context"

	justgrpcservice "github.com/9sarkan/golang-base/api/pb/github.com/9sarkan/golang-base/pkg/pb/justGRPCService.proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Controller struct {
	justgrpcservice.UnimplementedToDoListServer
}

func (Controller) Create(ctx context.Context, request *justgrpcservice.CreateToDoRequest) (*justgrpcservice.ToDoSignleResponse, error) {
	return nil, nil
}

func (Controller) Get(ctx context.Context, _ *emptypb.Empty) (*justgrpcservice.ToDoListResponse, error) {
	return nil, nil
}
