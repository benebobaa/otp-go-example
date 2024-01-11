package domain

import "time"

type UserRegister struct {
	Id         int
	Name       string
	Email      string
	Password   string
	IsVerified bool
	CreatedAt  time.Time
}
