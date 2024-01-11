package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"sent-email-otp/helper"
	"sent-email-otp/model/web"
	"sent-email-otp/model/web/user"
	"sent-email-otp/service"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (u *UserControllerImpl) CreateUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userCreateRequest := user.UserRequest{}

	helper.ReadFromRequestBody(request, &userCreateRequest)

	userResponse := u.UserService.CreateUser(request.Context(), userCreateRequest)

	webResponse := web.BaseResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
