package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type StatisticPostgres struct {
	db *sqlx.DB
}

func NewStatisticPostgres(db *sqlx.DB) *StatisticPostgres {
	return &StatisticPostgres{db: db}
}

func (r *StatisticPostgres) AddUser(userId int64, username string) error {
	const op = "postgres.AddUser"

	query := fmt.Sprintf("INSERT INTO %s (user_id, username) VALUES ($1, $2) ON CONFLICT (user_id) DO UPDATE SET username=EXCLUDED.username", usersTable)

	if _, err := r.db.Exec(query, userId, username); err != nil {
		return fmt.Errorf("%s%s", err, op)
	}

	return nil
}

func (r *StatisticPostgres) GetCountOfUsers() (int, error) {
	const op = "postgres.GetCountOfUsers"

	query := fmt.Sprintf("SELECT count(*) FROM %s", usersTable)

	var count int

	if err := r.db.QueryRow(query).Scan(&count); err != nil {
		return 0, fmt.Errorf("%s%s", err, op)
	}

	return count, nil
}