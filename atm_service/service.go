package atm_service

import "learn-cli/models"

type IService interface {
	Withdraw(number int, user *models.User) error
}

type Service struct {
}

func NewService() IService {
	return &Service{}
}

func (service *Service) Withdraw(number int, user *models.User) error {
	return nil
}
