package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type PengajuanUjianTaController interface {
	PengajuanUjianTa(writer http.ResponseWriter, r *http.Request, param httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
