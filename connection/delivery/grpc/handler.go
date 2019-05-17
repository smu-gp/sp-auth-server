package grpc

import (
	"context"
	_usecase "github.com/smu-gp/sp-sync-server/connection/usecase"
	connectionProtobuf "github.com/smu-gp/sp-sync-server/protobuf/connection"
	"google.golang.org/grpc"
	"io"
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
		accepted, err := server.usecase.RequestAuth(userId, req.DeviceInfo)
		if err != nil {
			return &connectionProtobuf.AuthResponse{Message: connectionProtobuf.AuthResponse_MESSAGE_FAILED}, err
		}
		if accepted {
			return &connectionProtobuf.AuthResponse{Message: connectionProtobuf.AuthResponse_MESSAGE_SUCCESS, UserId: userId}, nil
		} else {
			return &connectionProtobuf.AuthResponse{Message: connectionProtobuf.AuthResponse_MESSAGE_FAILED}, err
		}
	} else {
		return &connectionProtobuf.AuthResponse{Message: connectionProtobuf.AuthResponse_MESSAGE_FAILED}, nil
	}
}

func (server *server) WaitAuth(stream connectionProtobuf.ConnectionService_WaitAuthServer) error {
	req, err := stream.Recv()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}
	err = server.usecase.WaitAuth(req.UserId, stream)
	if err != nil {
		return err
	}
	return nil
}
