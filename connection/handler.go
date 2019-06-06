package connection

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
	response := &connectionProtobuf.AuthResponse{}
	userId, err := server.usecase.Auth(req.ConnectionCode)
	if err != nil {
		response.Message = connectionProtobuf.AuthResponse_MESSAGE_FAILED
		response.FailedReason = connectionProtobuf.AuthResponse_INTERNAL_ERR
		return response, err
	}
	if len(userId) > 0 {
		accepted, reason, err := server.usecase.RequestAuth(userId, req.DeviceInfo)
		if err != nil {
			response.Message = connectionProtobuf.AuthResponse_MESSAGE_FAILED
			response.FailedReason = reason
			return response, err
		}
		if accepted {
			response.Message = connectionProtobuf.AuthResponse_MESSAGE_SUCCESS
			response.UserId = userId
			return response, nil
		} else {
			response.Message = connectionProtobuf.AuthResponse_MESSAGE_FAILED
			response.FailedReason = reason
			return response, err
		}
	} else {
		response.Message = connectionProtobuf.AuthResponse_MESSAGE_FAILED
		response.FailedReason = connectionProtobuf.AuthResponse_AUTH_FAILED
		return response, nil
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
