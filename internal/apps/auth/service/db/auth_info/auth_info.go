package auth_info

import (
	"database/sql"
	"errors"
	"polybay/internal/apps/auth/service/db"
	"polybay/internal/apps/auth/service/model"
)

type authInfo struct {
	conn *sql.DB
}

func New(database model.Database) (model.IAuthInfo, error) {
	conn, err := db.NewConnection(database)
	if err != nil {
		return nil, err
	}
	return &authInfo{conn: conn}, nil
}

func (a *authInfo) CreateTableIfNotExists() error {
	return nil
}

func (a *authInfo) SelectById(id int64) (*model.AuthInfo, error) {
	return nil, errors.New("not implemented")
}

func (a *authInfo) Insert(info *model.AuthInfo) (*model.AuthInfo, error) {
	return nil, errors.New("not implemented")
}

func (a *authInfo) UpdatePassword(id int64, password string) (*model.AuthInfo, error) {
	return nil, errors.New("not implemented")
}

func (a *authInfo) UpdateEmail(id int64, email string) (*model.AuthInfo, error) {
	return nil, errors.New("not implemented")
}

func (a *authInfo) Delete(info *model.AuthInfo) error {
	return errors.New("not implemented")
}
