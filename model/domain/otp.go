package domain

import (
	"database/sql"
	"time"
)

type Otp struct {
	Id             int
	RefCode        string
	OtpValue       string
	ExpirationTime time.Time
	CreationTime   time.Time
	UserRegisId    sql.NullInt32
	UserUUID       sql.NullString
}
