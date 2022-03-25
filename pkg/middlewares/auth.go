package middlewares

import (
	"context"
	"errors"
	"strings"

	"github.com/9sarkan/golang-base/pkg/oauth"
	"google.golang.org/grpc/metadata"
)

func GRPCAuth(ctx context.Context) (context.Context, error) {
	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("401 unauthorized")
	}

	authorization := meta.Get("authorization")
	if len(authorization) == 0 {
		return nil, errors.New("401 unauthorized")
	}

	strArr := strings.Split(authorization[0], " ")

	if len(strArr) != 2 {
		return nil, errors.New("401 unauthorized")
	}

	resp, err := oauth.JWT.VerifyToken(ctx, strArr[1])
	if err != nil {
		return nil, err
	}

	ctx = context.WithValue(ctx, "user_id", resp.UserID) // the user uuid
	ctx = context.WithValue(ctx, "dialer", resp.Dialer)  // the user mobile

	return ctx, nil
}
