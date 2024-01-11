package otp

import (
	"database/sql"
	"time"
)

type OtpRequest struct {
	RefCode        string         `json:"ref_code"`
	OtpValue       string         `json:"otp_value"`
	ExpirationTime time.Time      `json:"expiration_time"`
	UserRegisId    sql.NullInt32  `json:"user_regis_id"`
	UserUUID       sql.NullString `json:"user_uuid"`
}

type OtpValidateRequest struct {
	RefCode  string `json:"ref_code" validate:"required"`
	OtpValue string `json:"otp_value" validate:"required,min=6,max=6"`
}
