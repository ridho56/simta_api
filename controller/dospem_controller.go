package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type DospemController interface {
	SetDospem(writer http.ResponseWriter, r *http.Request, param httprouter.Params)
	GetDospem(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetMhs(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetDosen(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
