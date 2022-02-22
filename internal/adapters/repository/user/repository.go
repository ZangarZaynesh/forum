package user

import (
	"context"
	"database/sql"
	"errors"

	"github.com/ZangarZaynesh/forum/internal/adapters/repository"
	"github.com/ZangarZaynesh/forum/internal/module"
)

type repo struct {
	db *sql.DB
}

// func NewRepository(db *sql.DB) (Create(dto *module.CreateUserDTO) error,Delete(id int) error) {
// 	return &repository{db: db}
// }

func NewRepository(db *sql.DB) repository.User {
	return &repo{db: db}
}

func (r *repo) CheckByLogin(ctx context.Context, dto *module.CreateUserDTO) error {
	if !test("login", dto.Login, r) {
		return errors.New("This login already exists")
	}
	return nil
}

func (r *repo) CheckByEmail(ctx context.Context, dto *module.CreateUserDTO) error {
	if !test("email", dto.Email, r) {
		return errors.New("This email already exists")
	}
	return nil
}

func (r *repo) Create(ctx context.Context, dto *module.CreateUserDTO) error {
	_, err := r.db.Exec("INSERT INTO users (login, password, email) VALUES ( ?, ?, ?);", dto.Login, dto.Password, dto.Email)
	if err != nil {
		return errors.New("400 Bad Request")
	}
	return nil
}

func (r *repo) Delete(id int) error {
	return nil
}

func test(NameColumn, ValueColumn string, r *repo) bool {
	row := r.db.QueryRow("SELECT ? FROM users WHERE users."+NameColumn+"= ?;", NameColumn, ValueColumn)
	err := row.Scan()
	return errors.Is(err, sql.ErrNoRows)
}
