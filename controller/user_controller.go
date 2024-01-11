package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type UserController interface {
	CreateUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
