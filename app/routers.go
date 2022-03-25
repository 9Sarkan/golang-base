package app

import (
	"context"
	"fmt"
	"log"
	"net"

	justgrpcservice "github.com/9sarkan/golang-base/api/pb/github.com/9sarkan/golang-base/pkg/pb/justGRPCService.proto"
	"github.com/9sarkan/golang-base/config"
	"github.com/9sarkan/golang-base/controller/todo"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_middleware_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type RouterCenter struct{}

// InitialGrpcServer initial a grpc server with auth, prometheus and panic recover
func (RouterCenter) InitialGrpcServer(ctx context.Context, cnf config.HTTPServer, authFunc grpc_middleware_auth.AuthFunc,
	panicRecoverFunc grpc_recovery.RecoveryHandlerFuncContext) error {

	listener, err := net.Listen(cnf.Protocol, fmt.Sprintf("%s:%d", cnf.Host, cnf.Port))
	if err != nil {
		log.Default().Printf(err.Error())
		return err
	}

	// create options for grpc server
	options := []grpc.ServerOption{
		// Unary interceptors
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_middleware_auth.UnaryServerInterceptor(authFunc),
				grpc_prometheus.UnaryServerInterceptor,
				grpc_recovery.UnaryServerInterceptor(grpc_recovery.WithRecoveryHandlerContext(panicRecoverFunc)),
			),
		),
	}

	baseServer := grpc.NewServer(options...)

	justgrpcservice.RegisterToDoListServer(baseServer, &todo.Controller{})

	reflection.Register(baseServer)

	// log if server down
	go func() {
		log.Default().Fatalln(baseServer.Serve(listener))
	}()

	return nil
}

func (RouterCenter) InitialHttpRouter(cnf config.HTTPServer) {}

func (RouterCenter) InitialGatewayRouter(cnf config.HTTPServer) {}
