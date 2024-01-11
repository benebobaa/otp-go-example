package otp

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"sent-email-otp/helper"
	"sent-email-otp/model/domain"
)

type OtpRepositoryImpl struct{}

func NewOtpRepository() OtpRepository {
	return &OtpRepositoryImpl{}
}

func (o *OtpRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, otp domain.Otp) domain.Otp {

	//if otp.UserRegisId != 0 {
	//	sqlRegis := "INSERT INTO otp_users (ref_code, otp_value, expiration_time, user_regis_id) VALUES ($1, $2, $3, $4) RETURNING ref_code"
	//	rowRegis := tx.QueryRowContext(ctx, sqlRegis, otp.RefCode, otp.OtpValue, otp.ExpirationTime, otp.UserRegisId)
	//	err := rowRegis.Scan(&otp.RefCode)
	//	helper.PanicIfError(err)
	//
	//	return otp
	//}

	SQL := "INSERT INTO otp_users (ref_code, otp_value, expiration_time,user_regis_id,user_uuid) VALUES ($1, $2, $3, $4, $5) RETURNING ref_code"

	row := tx.QueryRowContext(ctx, SQL, otp.RefCode, otp.OtpValue, otp.ExpirationTime, otp.UserRegisId, otp.UserUUID)

	err := row.Scan(&otp.RefCode)

	helper.PanicIfError(err)

	return otp
}

func (o *OtpRepositoryImpl) FindByRefCode(ctx context.Context, tx *sql.Tx, refCode string) (domain.Otp, error) {
	fmt.Println("refCode", refCode)
	SQL := "SELECT id, ref_code, otp_value, expiration_time, creation_time, user_regis_id, user_uuid FROM otp_users WHERE ref_code = $1 AND expiration_time > NOW()"

	rows, err := tx.QueryContext(ctx, SQL, refCode)
	fmt.Println("rows", rows)
	helper.PanicIfError(err)
	defer rows.Close()

	otp := domain.Otp{}

	if rows.Next() {
		fmt.Println("ERROR??")
		errScan := rows.Scan(&otp.Id, &otp.RefCode, &otp.OtpValue, &otp.ExpirationTime, &otp.CreationTime, &otp.UserRegisId, &otp.UserUUID)
		fmt.Println("errScan", errScan)
		helper.PanicIfError(errScan)
		fmt.Println("otp", otp)
		return otp, nil
	} else {
		fmt.Println("otp", otp)
		return otp, errors.New("otp not found or expired")
	}
}
