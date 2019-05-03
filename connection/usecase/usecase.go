package usecase

import "github.com/smu-gp/sp-sync-server/connection/repository"

func NewConnectionUsecase(repository repository.ConnectionRepository) ConnectionUsecase {
	return &connectionUsecase{repository}
}

type ConnectionUsecase interface {
	RequestUserId()
	InitConnection(userId string) string
	Auth(connectionCode string) (string, error)
}

type connectionUsecase struct {
	connectionRepository repository.ConnectionRepository
}

func (usecase *connectionUsecase) RequestUserId() {
	panic("implement me")
}

func (usecase *connectionUsecase) InitConnection(userId string) string {
	panic("implement me")
}

func (usecase *connectionUsecase) Auth(connectionCode string) (string, error) {
	panic("implement me")
}


