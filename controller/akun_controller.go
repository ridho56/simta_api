package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type AkunController interface {
	//Register(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Login(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	//EmailCheck(writer http.ResponseWriter, r *http.Request, params httprouter.Params)
}
