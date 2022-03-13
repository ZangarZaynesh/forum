package user

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"github.com/ZangarZaynesh/forum/internal/adapters/repository"
	"github.com/ZangarZaynesh/forum/internal/module"

	// "go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
	"golang.org/x/crypto/bcrypt"
)

type repo struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) repository.User {
	return &repo{db: db}
}

func (r *repo) CheckByLogin(ctx context.Context, dto *module.CreateUserDTO) error {
	if !test("login", dto.Login, r) {
		return errors.New("This login already exists")
	}
	return nil
}

func (r *repo) CheckSignIn(ctx context.Context, dto *module.SignUserDTO) error {
	var password []byte
	row := r.db.QueryRow("SELECT id ,password FROM users WHERE users.login= ?;", dto.Login)
	err := row.Scan(&dto.UserId, &password)
	if errors.Is(err, sql.ErrNoRows) {
		return errors.New("This login does not exist")
	}

	err = bcrypt.CompareHashAndPassword(password, []byte(dto.Password))
	if err != nil {
		return errors.New("Invalid password")
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

func (r *repo) AddCookie(ctx context.Context, dto *module.SignUserDTO) error {
	_, err := r.db.Exec("INSERT INTO sessions (user_id, key, date, duration) VALUES ( ?, ?, ?, ?);", dto.UserId, dto.UUID, dto.CreateTimeUUID, dto.Duration)
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) CheckCookie(ctx context.Context, session *http.Cookie, dto *module.HomePageDTO) error {
	row := r.db.QueryRow("SELECT user_id FROM sessions where key = ? ;", session.Value)
	err := row.Scan(&dto.UserId)
	if errors.Is(err, sql.ErrNoRows) {
		return err
	}
	return nil
}

func (r *repo) DeleteSession(ctx context.Context, userId int) error {
	_, err := r.db.Exec("DELETE FROM sessions WHERE user_id = ?;", userId)
	if err != nil {
		return err
	}

	return nil
}

func test(NameColumn, ValueColumn string, r *repo) bool {
	row := r.db.QueryRow("SELECT ? FROM users WHERE users."+NameColumn+"= ?;", NameColumn, ValueColumn)
	err := row.Scan()
	return errors.Is(err, sql.ErrNoRows)
}
