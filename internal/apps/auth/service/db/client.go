package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"polybay/internal/apps/auth/service/model"
)

func NewConnection(database model.Database) (*sql.DB, error) {
	connStr := fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%s sslmode=disable",
		database.Name, database.User, database.Password, database.Ip, database.Port)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}
