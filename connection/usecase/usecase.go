package usecase

import (
	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"github.com/pquerna/otp/totp"
	"github.com/smu-gp/sp-sync-server/connection/repository"
	connectionGrpc "github.com/smu-gp/sp-sync-server/protobuf/build"
	"strconv"
	"strings"
	"time"
)

func NewConnectionUsecase(repository repository.ConnectionRepository) ConnectionUsecase {
	return &connectionUsecase{repository}
}

type ConnectionUsecase interface {
	RequestUserId() (string, error)
	Connection(userId string) (string, error)
	Auth(connectionCode string) (string, error)
	RequestAuth(userId string, deviceInfo *connectionGrpc.AuthDeviceInfo) (bool, error)
	WaitAuth(userId string, stream connectionGrpc.ConnectionService_WaitAuthServer) error
	ResponseAuth(userId string, accept bool) error
}

type connectionUsecase struct {
	repository repository.ConnectionRepository
}

func (usecase *connectionUsecase) RequestUserId() (string, error) {
	return uuid.New().String(), nil
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

func (usecase *connectionUsecase) RequestAuth(userId string, deviceInfo *connectionGrpc.AuthDeviceInfo) (bool, error) {
	var message = deviceInfo.DeviceName + ","
	if deviceInfo.DeviceType == connectionGrpc.AuthDeviceInfo_DEVICE_TABLET {
		message += "TABLET"
	} else {
		message += "WEB"
	}
	pubSub := usecase.repository.Subscribe("auth_res:" + userId)
	defer pubSub.Close()

	err := usecase.repository.Publish("auth:"+userId, message)
	if err != nil {
		return false, err
	}

	for {
		iface, err := pubSub.Receive()
		if err != nil {
			return false, err
		}

		switch msg := iface.(type) {
		case *redis.Message:
			accepted, err := strconv.ParseBool(msg.Payload)
			if err != nil {
				return false, err
			}
			return accepted, nil
		}
	}
}

func (usecase *connectionUsecase) WaitAuth(userId string, stream connectionGrpc.ConnectionService_WaitAuthServer) error {
	pubSub := usecase.repository.Subscribe("auth:" + userId)
	defer pubSub.Close()

	for {
		iface, err := pubSub.Receive()
		if err != nil {
			return err
		}

		switch msg := iface.(type) {
		case *redis.Message:
			var deviceMessage = strings.Split(msg.Payload, ",")
			var deviceType connectionGrpc.AuthDeviceInfo_DeviceType
			if strings.Compare(deviceMessage[1], "TABLET") == 0 {
				deviceType = connectionGrpc.AuthDeviceInfo_DEVICE_TABLET
			} else {
				deviceType = connectionGrpc.AuthDeviceInfo_DEVICE_WEB
			}
			err = stream.Send(&connectionGrpc.WaitAuthResponse{
				AuthDevice: &connectionGrpc.AuthDeviceInfo{
					DeviceName: deviceMessage[0],
					DeviceType: deviceType,
				},
			})
			if err != nil {
				return err
			}
			req, err := stream.Recv()
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
	}
}

func (usecase *connectionUsecase) ResponseAuth(userId string, accept bool) error {
	return usecase.repository.Publish("auth_res:"+userId, strconv.FormatBool(accept))
}
