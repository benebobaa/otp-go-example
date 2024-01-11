package repository

import (
	"context"
	"database/sql"
	"sent-email-otp/helper"
	"sent-email-otp/model/domain"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (u *UserRepositoryImpl) CreateUser(ctx context.Context, tx *sql.DB, user domain.User) domain.User {
	SQL := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at"

	row := tx.QueryRowContext(ctx, SQL, user.Name, user.Email, user.Password)

	err := row.Scan(&user.Id, &user.CreatedAt, &user.UpdatedAt)
	helper.PanicIfError(err)

	return user
}
