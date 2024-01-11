package main

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
	"net/http"
	"sent-email-otp/app"
	"sent-email-otp/controller"
	"sent-email-otp/helper"
	"sent-email-otp/repository"
	"sent-email-otp/repository/otp"
	"sent-email-otp/service"
	otpServ "sent-email-otp/service/otp"
)

func main() {
	db := app.NewDB()
	validate := validator.New()

	//USER
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)

	//OTP
	otpRepository := otp.NewOtpRepository()
	otpService := otpServ.NewOtpService(otpRepository, db, validate)
	otpController := controller.NewOtpController(otpService)

	//AUTH
	authRepository := repository.NewAuthRepository()
	authService := service.NewAuthService(authRepository, otpRepository, db, validate)
	authController := controller.NewAuthController(authService)

	router := app.NewRouter(userController, authController, otpController)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
