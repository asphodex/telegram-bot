package service

import "telegram-bot/internal/telegram-bot/repository"

type StatisticService struct {
	repo repository.Statistic
}

func NewStatisticService(repo repository.Statistic) *StatisticService {
	return &StatisticService{repo: repo}
}

func (s *StatisticService) AddUser(userId int64, userName string) error {
	return s.repo.AddUser(userId, userName)
}

func (s *StatisticService) GetCountOfUsers() (int, error) {
	return s.repo.GetCountOfUsers()
}
