package repository

import (
	"context"
	"database/sql"
	"sent-email-otp/helper"
	"sent-email-otp/model/domain"
)

type AuthRepositoryImpl struct{}

func NewAuthRepository() AuthRepository {
	return &AuthRepositoryImpl{}
}

func (a *AuthRepositoryImpl) Login(ctx context.Context, tx *sql.Tx, user domain.User) {
	//TODO implement me
	panic("implement me")
}

func (a *AuthRepositoryImpl) Register(ctx context.Context, tx *sql.Tx, userRegister domain.UserRegister) domain.UserRegister {
	SQL := "INSERT INTO user_registrations (name, email, password) VALUES ($1, $2, $3) RETURNING id"

	row := tx.QueryRowContext(ctx, SQL, userRegister.Name, userRegister.Email, userRegister.Password)

	err := row.Scan(&userRegister.Id)

	helper.PanicIfError(err)

	return userRegister
}
