package db

import (
	"context"
	"hello/internals/app/models"
	"time"

	"github.com/jackc/pgx/v5"
	log "github.com/sirupsen/logrus"
)

type UsersStorage struct {
	databasePool *pgx.Conn
}

func NewUsersStorage(pool *pgx.Conn) *UsersStorage {
	storage := new(UsersStorage)
	storage.databasePool = pool
	return storage
}

func (storage *UsersStorage) CreateUser(user models.User) error {
	query := "insert INTO users(username, email, created) VALUES ($1, $2, $3)"

	user.Created = time.Now().Format("2006-01-02 15:04:05")
	_, err := storage.databasePool.Exec(context.Background(), query, user.Username, user.Email, user.Created)

	if err != nil {
		log.Errorln(err)
		return err
	}

	return nil
}
