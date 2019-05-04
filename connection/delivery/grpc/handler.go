package grpc

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	_usecase "github.com/smu-gp/sp-sync-server/connection/usecase"
	connectionProtobuf "github.com/smu-gp/sp-sync-server/protobuf/connection"
	"google.golang.org/grpc"
)

func NewConnectionGrpcServer(grpcServer *grpc.Server, connectionUsecase _usecase.ConnectionUsecase) {
	connectionServer := &server{
		usecase: connectionUsecase,
	}
	connectionProtobuf.RegisterConnectionServiceServer(grpcServer, connectionServer)
}

type server struct {
	usecase _usecase.ConnectionUsecase
}

func (server *server) RequestUserId(context.Context, *empty.Empty) (*connectionProtobuf.RequestUserIdResponse, error) {
	userId, err := server.usecase.RequestUserId()
	return &connectionProtobuf.RequestUserIdResponse{UserId: userId}, err
}

func (server *server) Connection(ctx context.Context, req *connectionProtobuf.ConnectionRequest) (*connectionProtobuf.ConnectionResponse, error) {
	code, err := server.usecase.Connection(req.UserId)
	return &connectionProtobuf.ConnectionResponse{ConnectionCode: code}, err
}

func (server *server) Auth(ctx context.Context, req *connectionProtobuf.AuthRequest) (*connectionProtobuf.AuthResponse, error) {
	userId, err := server.usecase.Auth(req.ConnectionCode)
	if err != nil {
		return &connectionProtobuf.AuthResponse{Message: connectionProtobuf.AuthResponse_MESSAGE_FAILED}, err
	}
	if len(userId) > 0 {
		return &connectionProtobuf.AuthResponse{Message: connectionProtobuf.AuthResponse_MESSAGE_SUCCESS, UserId: userId}, nil
	} else {
		return &connectionProtobuf.AuthResponse{Message: connectionProtobuf.AuthResponse_MESSAGE_FAILED}, nil
	}
}

