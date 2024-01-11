package repository

import (
	"context"
	"database/sql"
	"sent-email-otp/model/domain"
)

type UserRepository interface {
	CreateUser(ctx context.Context, tx *sql.DB, user domain.User) domain.User
}
