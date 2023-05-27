package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type PengajuanSeminarHasilController interface {
	UploadFile(writer http.ResponseWriter, r *http.Request, param httprouter.Params)
	GetFile(writer http.ResponseWriter, r *http.Request, param httprouter.Params)
	PengajuanJadwalSemhas(writer http.ResponseWriter, r *http.Request, param httprouter.Params)
	UpdateAjuanSemhas(writer http.ResponseWriter, r *http.Request, param httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
