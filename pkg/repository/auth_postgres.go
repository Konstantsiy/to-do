package repository

import (
	"fmt"
	"github.com/Konstantsiy/todo/pkg/entity"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user entity.User) (int, error) {
	var id int
	query := fmt.Sprintf("insert into %s (name, username, password_hash) values ($1, $2, $3) returning id", userTable)
	row := r.db.QueryRowx(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return 0, nil
}

func (r *AuthPostgres) GetUser(username, password string) (entity.User, error) {
	var user entity.User
	query := fmt.Sprintf("select id from %s where username=$1 and password_hash=$2", userTable)
	err := r.db.Get(&user, query, username, password)
	return user, err
}
