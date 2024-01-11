package app

import (
	"github.com/julienschmidt/httprouter"
	"sent-email-otp/controller"
	"sent-email-otp/exception"
)

func NewRouter(userController controller.UserController, authController controller.AuthController, otpController controller.OtpController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/v1/users", userController.CreateUser)
	//router.POST("/api/v1/auth/login", authController.Login)
	router.POST("/api/v1/auth/register", authController.Register)

	router.POST("/api/v1/otp/validate", otpController.FindAndValidate)

	router.PanicHandler = exception.ErrorHandler

	return router
}
