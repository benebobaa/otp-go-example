package otp

import (
	"context"
	"database/sql"
	"sent-email-otp/model/domain"
)

type OtpRepository interface {
	Create(ctx context.Context, tx *sql.Tx, otp domain.Otp) domain.Otp
	FindByRefCode(ctx context.Context, tx *sql.Tx, refCode string) (domain.Otp, error)
}
