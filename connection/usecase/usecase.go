package usecase

import (
	"github.com/google/uuid"
	"github.com/pquerna/otp/totp"
	"github.com/smu-gp/sp-sync-server/connection/repository"
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
}

type connectionUsecase struct {
	connectionRepository repository.ConnectionRepository
}

func (usecase *connectionUsecase) RequestUserId() (string, error) {
	return uuid.New().String(), nil
}

func (usecase *connectionUsecase) Connection(userId string) (string, error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer: "smu-gp",
		AccountName: userId,
	})
	if key != nil {
		_, err := usecase.connectionRepository.StoreSecret(userId, key.Secret())
		code, err := totp.GenerateCode(key.Secret(), time.Now())
		return code, err
	} else {
		return "", err
	}
}

func (usecase *connectionUsecase) Auth(connectionCode string) (string, error) {
	keys, err := usecase.connectionRepository.GetAllConnection()
	if err != nil {
		return "", err
	}
	for i := range keys {
		var key = keys[i]
		secret, err := usecase.connectionRepository.GetSecret(key)
		if err != nil {
			return "", err
		}
		if totp.Validate(connectionCode, *secret) {
			return strings.Split(key, ":")[1], nil
		}
	}
	return "", nil
}
