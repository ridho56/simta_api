package app

import (
	"github.com/julienschmidt/httprouter"
	"goelster/controller"
	"goelster/exception"
)

func NewRouter(
	controllerAkun controller.AkunController,
	controllerPengajuan controller.PengajuanBimbinganController,
	controllerSempro controller.PengajuanSemproController,
	controllerDospem controller.DospemController,
	controllerSeminarHasil controller.PengajuanSeminarHasilController,
	controllerUjianTa controller.PengajuanUjianTaController,
) *httprouter.Router {
	router := httprouter.New()
	//router.POST("/simta/register", controllerAkun.Register)
	router.POST("/simta/login", controllerAkun.Login)
	//router.GET("/simta/emailcheck/:emailUser", controllerAkun.EmailCheck)
	router.POST("/simta/bimbingan", controllerPengajuan.PengajuanBimbingan)
	router.GET("/simta/bimbinganid/:Id", controllerPengajuan.FindById)
	router.GET("/simta/bimbingandosen/:Id", controllerPengajuan.FindByIdDosen)
	router.PUT("/simta/bimbinganupdate", controllerPengajuan.UpdateDataAjuanBimbingan)
	router.PUT("/simta/hasilbim", controllerPengajuan.Hasilbim)
	router.POST("/simta/sempro", controllerSempro.Pengajuansempro)
	router.PUT("/simta/semproupdate", controllerSempro.UpdateAjuanSempro)
	router.GET("/simta/semproid/:Id", controllerSempro.FindById)
	router.POST("/simta/setdospem", controllerDospem.SetDospem)
	router.GET("/simta/getdospem/:Id", controllerDospem.GetDospem)
	router.GET("/simta/getmhs", controllerDospem.GetMhs)
	router.GET("/simta/getdosen", controllerDospem.GetDosen)
	router.POST("/simta/seminarhasil", controllerSeminarHasil.UploadFile)
	router.GET("/simta/getseminar", controllerSeminarHasil.GetFile)
	router.POST("/simta/jadwalsemhas", controllerSeminarHasil.PengajuanJadwalSemhas)
	router.PUT("/simta/updatesemhas", controllerSeminarHasil.UpdateAjuanSemhas)
	router.GET("/simta/semhasid/:Id", controllerSeminarHasil.FindById)
	router.POST("/simta/pengajuanta", controllerUjianTa.PengajuanUjianTa)
	router.GET("/simta/ujiantaid/:Id", controllerUjianTa.FindById)
	router.PanicHandler = exception.ErrorHandler

	return router

}
