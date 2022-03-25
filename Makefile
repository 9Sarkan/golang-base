protogrpc:
	protoc --go_out=api/pb --go-grpc_out=api/pb pkg/pb/grpc/*.proto