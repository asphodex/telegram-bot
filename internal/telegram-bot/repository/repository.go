package repository

import (
	"github.com/jmoiron/sqlx"
	"telegram-bot/internal/telegram-bot/repository/postgres"
)

type Statistic interface {
	GetCountOfUsers() (int, error)
	AddUser(userId int64, username string) error
}

type Repository struct {
	Statistic
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Statistic:postgres.NewStatisticPostgres(db),
	}
}
