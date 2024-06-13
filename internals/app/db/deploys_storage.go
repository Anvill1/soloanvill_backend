package db

import (
	"context"
	"hello/internals/app/models"
	"time"

	"github.com/jackc/pgx/v5"
	log "github.com/sirupsen/logrus"
)

type DeployStorage struct {
	databasePool *pgx.Conn
}

func NewDeployStorage(pool *pgx.Conn) *DeployStorage {
	storage := new(DeployStorage)
	storage.databasePool = pool
	return storage
}

func (storage *DeployStorage) CreateDeploy(user models.User, clientIP string) error {
	query := "insert INTO users(username, email, created, ip) VALUES ($1, $2, $3, $4)"

	user.Created = time.Now().Format("2006-01-02 15:04:05")
	_, err := storage.databasePool.Exec(context.Background(), query, user.Username, user.Email, user.Created, clientIP)

	if err != nil {
		log.Errorln(err)
		return err
	}

	return nil
}
