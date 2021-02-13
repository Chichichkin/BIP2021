package model

import "time"

type IAuthInfo interface {
	CreateTableIfNotExists() error
	SelectById(id int64) (*AuthInfo, error)
	Insert(info *AuthInfo) (*AuthInfo, error)
	UpdatePassword(id int64, password string) (*AuthInfo, error)
	UpdateEmail(id int64, email string) (*AuthInfo, error)
	Delete(info *AuthInfo) error
}

type AuthInfo struct {
	Id        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Role      Role      `json:"role"`
	Status    Status    `json:"status"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
}

type Role int

const (
	User Role = iota
	Moderator
)

func (r Role) String() string {
	return [...]string{"User", "Moderator"}[r]
}

type Status int

const (
	NotConfirmed Status = iota
	Confirmed
	Banned
)

func (s Status) String() string {
	return [...]string{"NotConfirmed", "Confirmed", "Banned"}[s]
}
