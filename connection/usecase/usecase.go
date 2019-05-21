package usecase

import (
	"github.com/golang/protobuf/proto"
	"github.com/pquerna/otp/totp"
	"github.com/smu-gp/sp-sync-server/connection/repository"
	connectionGrpc "github.com/smu-gp/sp-sync-server/protobuf/connection"
	"io"
	"strconv"
	"strings"
	"time"
)

func NewConnectionUsecase(repository repository.ConnectionRepository) ConnectionUsecase {
	return &connectionUsecase{repository}
}

type ConnectionUsecase interface {
	Connection(userId string) (string, error)
	Auth(connectionCode string) (string, error)
	RequestAuth(userId string, deviceInfo *connectionGrpc.AuthDeviceInfo) (bool, connectionGrpc.AuthResponse_FailedReason, error)
	WaitAuth(userId string, stream connectionGrpc.ConnectionService_WaitAuthServer) error
	ResponseAuth(userId string, accept bool) error
}

type connectionUsecase struct {
	repository repository.ConnectionRepository
}

func (usecase *connectionUsecase) Connection(userId string) (string, error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "smu-gp",
		AccountName: userId,
	})
	if key != nil {
		_, err := usecase.repository.StoreSecret(userId, key.Secret())
		code, err := totp.GenerateCode(key.Secret(), time.Now())
		return code, err
	} else {
		return "", err
	}
}

func (usecase *connectionUsecase) Auth(connectionCode string) (string, error) {
	keys, err := usecase.repository.GetAllConnection()
	if err != nil {
		return "", err
	}
	for i := range keys {
		var key = keys[i]
		secret, err := usecase.repository.GetSecret(key)
		if err != nil {
			return "", err
		}
		if totp.Validate(connectionCode, *secret) {
			deleted, err := usecase.repository.DeleteKey(key)
			if err != nil || !deleted {
				return "", err
			}
			return strings.Split(key, ":")[1], nil
		}
	}
	return "", nil
}

func (usecase *connectionUsecase) RequestAuth(userId string, deviceInfo *connectionGrpc.AuthDeviceInfo) (bool, connectionGrpc.AuthResponse_FailedReason, error) {
	pubSub := usecase.repository.Subscribe("auth_res:" + userId)
	defer pubSub.Close()

	deviceInfoData, _ := proto.Marshal(deviceInfo)
	subs, err := usecase.repository.Publish("auth:"+userId, string(deviceInfoData))
	if err != nil {
		return false, connectionGrpc.AuthResponse_INTERNAL_ERR, err
	}
	if subs == 0 {
		return false, connectionGrpc.AuthResponse_NO_HOST_WAITED, nil
	}

	ch := pubSub.Channel()

	time.AfterFunc(time.Minute, func() {
		_ = pubSub.Close()
	})

	for msg := range ch {
		accepted, _ := strconv.ParseBool(msg.Payload)
		var reason = connectionGrpc.AuthResponse_REJECT_HOST
		if accepted {
			reason = connectionGrpc.AuthResponse_NONE
		}
		return accepted, reason, nil
	}
	return false, connectionGrpc.AuthResponse_RESPONSE_TIMEOUT, nil
}

func (usecase *connectionUsecase) WaitAuth(userId string, stream connectionGrpc.ConnectionService_WaitAuthServer) error {
	pubSub := usecase.repository.Subscribe("auth:" + userId)
	defer pubSub.Close()

	ch := pubSub.Channel()

	for msg := range ch {
		var deviceInfo = &connectionGrpc.AuthDeviceInfo{}
		_ = proto.Unmarshal([]byte(msg.Payload), deviceInfo)
		err := stream.Send(&connectionGrpc.WaitAuthResponse{
			AuthDevice: deviceInfo,
		})
		if err != nil {
			return err
		}
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		err = usecase.ResponseAuth(userId, req.AcceptDevice)
		if err != nil {
			return err
		}
		if req.AcceptDevice {
			return nil
		}
	}
	return nil
}

func (usecase *connectionUsecase) ResponseAuth(userId string, accept bool) error {
	_, err := usecase.repository.Publish("auth_res:"+userId, strconv.FormatBool(accept))
	return err
}
