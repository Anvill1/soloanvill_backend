package db

import (
	"context"

	"os"

	"github.com/jackc/pgx/v5"
	log "github.com/sirupsen/logrus"
)

type DBStorage struct {
	conn *pgx.Conn
}

func NewDBStorage(conn *pgx.Conn) *DBStorage {
	storage := new(DBStorage)
	storage.conn = conn
	return storage
}

func (storage *DBStorage) CheckDB(table_name string) bool {
	var exists bool

	query := "SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_schema = 'public' AND table_name = $1)"

	err := storage.conn.QueryRow(context.Background(), query, table_name).Scan(&exists)
	if err != nil {
		log.Errorln(err)
	}
	return exists
}

func (storage *DBStorage) InitDB(migrationFile string) error {

	migrations, err := os.ReadFile(migrationFile)
	if err != nil {
		log.Errorln(err)
	}

	_, err = storage.conn.Exec(context.Background(), string(migrations))
	if err != nil {
		log.Errorln(err)
	}

	return nil
}
