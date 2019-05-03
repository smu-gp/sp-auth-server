package grpc

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/smu-gp/sp-sync-server/connection/delivery/grpc/connection_grpc"
	_usecase "github.com/smu-gp/sp-sync-server/connection/usecase"
	"google.golang.org/grpc"
)

func NewConnectionGrpcServer(grpcServer *grpc.Server, usecase _usecase.ConnectionUsecase) {
	connectionServer := &server{}
	connection_grpc.RegisterConnectionServiceServer(grpcServer, connectionServer)
}

type server struct {
}

func (server *server) RequestUserId(context.Context, *empty.Empty) (*connection_grpc.RequestUserIdResponse, error) {
	panic("implement me")
}

func (server *server) Connection(context.Context, *connection_grpc.ConnectionRequest) (*connection_grpc.ConnectionResponse, error) {
	panic("implement me")
}

func (server *server) Auth(context.Context, *connection_grpc.AuthRequest) (*connection_grpc.AuthResponse, error) {
	panic("implement me")
}
