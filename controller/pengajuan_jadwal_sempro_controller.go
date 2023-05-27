package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type PengajuanSemproController interface {
	Pengajuansempro(writer http.ResponseWriter, r *http.Request, param httprouter.Params)
	UpdateAjuanSempro(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
