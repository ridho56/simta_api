package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type PengajuanBimbinganController interface {
	PengajuanBimbingan(writer http.ResponseWriter, r *http.Request, param httprouter.Params)
	UpdateDataAjuanBimbingan(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindByIdDosen(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Hasilbim(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
