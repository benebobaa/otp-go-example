package repository

import (
	"context"
	"database/sql"
	"sent-email-otp/model/domain"
)

type AuthRepository interface {
	Login(ctx context.Context, tx *sql.Tx, user domain.User)
	Register(ctx context.Context, tx *sql.Tx, user domain.UserRegister) domain.UserRegister
}
