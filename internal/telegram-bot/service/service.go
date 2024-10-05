package service

import (
	"telegram-bot/internal/telegram-bot/repository"
)

type Statistic interface {
	GetCountOfUsers() (int, error)
	AddUser(userId int64, username string) error
}

type Service struct {
	Statistic
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Statistic: NewStatisticService(repo.Statistic),
	}
}