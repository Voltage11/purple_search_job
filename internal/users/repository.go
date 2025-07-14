package users

import (
	"context"
	"fmt"
	"lesson/pkg/utils"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type UserRepository struct {
	DbPool *pgxpool.Pool
	CustomLogger *zerolog.Logger
}

func NewUserRepository(dbPool *pgxpool.Pool, customLogger *zerolog.Logger) *UserRepository{
	return &UserRepository{
		DbPool: dbPool,
		CustomLogger: customLogger,
	}
}

func (r *UserRepository) GetByEmail(email string) (User, error) {
	query := "SELECT * FROM users WHERE email=@email"
	args := pgx.NamedArgs{"email": email,}
	row, err := r.DbPool.Query(context.Background(), query, args)
	if (err != nil) {
		return User{}, err
	}
	user, err := pgx.CollectOneRow(row, pgx.RowToStructByName[User])
	if (err != nil) {
		return User{}, err
	}
	return user, nil
}


func (r *UserRepository) AddUser(form UserCreateForm) (string, error) {
	query := `INSERT INTO users (name, email, pass_hash, createdat) VALUES (@name, @email, @pass_hash, @createdat)`
	args := pgx.NamedArgs{
		"name": form.Name,
		"email": form.Email,
		"pass_hash": utils.StrToHash(form.PassHash),
		"createdat": time.Now(),
	}
	_, err := r.DbPool.Exec(context.Background(), query, args)
	if (err != nil) {
		return "", fmt.Errorf("невозможно создать пользователя: %w", err)
	}
	return form.Email, nil
}