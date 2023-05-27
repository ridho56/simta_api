//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
	"goelster/app"
	"goelster/controller"
	"goelster/middleware"
	"goelster/repository"
	"goelster/service"
	"net/http"
)

var registerSet = wire.NewSet(
	repository.NewAkunRepositoryImpl,
	wire.Bind(new(repository.AkunRepository), new(*repository.AkunRepositoryImpl)),
	service.NewAkunService,
	wire.Bind(new(service.AkunService), new(*service.AkunServiceImpl)),
	controller.NewAkunController,
	wire.Bind(new(controller.AkunController), new(*controller.AkunControllerImpl)),
)
var bimbinganSet = wire.NewSet(
	repository.NewPengajuanBimbinganRepository,
	wire.Bind(new(repository.PengajuanBimbinganRepository), new(*repository.PengajuanBimbinganRepositoryImpl)),
	service.NewPegajuanBimbinganService,
	wire.Bind(new(service.PengajuanBimbinganService), new(*service.PengajuanBimbinganServiceImpl)),
	controller.NewPengajuanBimbinganController,
	wire.Bind(new(controller.PengajuanBimbinganController), new(*controller.PengajuanBimbinganControllerImpl)),
)
var semproSet = wire.NewSet(
	repository.NewPengajuanSemproRepository,
	wire.Bind(new(repository.PengajuanSemproRepository), new(*repository.PengajuanSemproRepositoryImpl)),
	service.NewPengajuanSemproService,
	wire.Bind(new(service.PengajuanSemproService), new(*service.PengajuanSemproServiceImpl)),
	controller.NewPengajuanSemproController,
	wire.Bind(new(controller.PengajuanSemproController), new(*controller.PengajuanSemproControllerImpl)),
)
var dospemset = wire.NewSet(
	repository.NewDospemRepository,
	wire.Bind(new(repository.DospemRepository), new(*repository.DospemRepositoryImpl)),
	service.NewDospemService,
	wire.Bind(new(service.DospemService), new(*service.DospemServiceImpl)),
	controller.NewDospemController,
	wire.Bind(new(controller.DospemController), new(*controller.DospemControllerImpl)),
)
var seminarhasilset = wire.NewSet(
	repository.NewPengajuanSeminarHasilRepository,
	wire.Bind(new(repository.PengajuanSeminarHasilRepository), new(*repository.PengajuanSeminarHasilRepositoryImpl)),
	service.NewPengajuanSeminarHasilService,
	wire.Bind(new(service.PengajuanSeminarHasilService), new(*service.PengajuanSeminarHasilServiceImpl)),
	controller.NewPengajuanSeminarHasilController,
	wire.Bind(new(controller.PengajuanSeminarHasilController), new(*controller.PengajuanSeminarHasilControllerImpl)),
)

var ujiantaset = wire.NewSet(
	repository.NewPengajuanUjianTaRepository,
	wire.Bind(new(repository.PengajuanUjianTaRepository), new(*repository.PengajuanUjianTaRepositoryImpl)),
	service.NewPengajuanUjianTaService,
	wire.Bind(new(service.PengajuanUjianTaService), new(*service.PengajuanUjianTaServiceImpl)),
	controller.NewPengajuanUjianTaController,
	wire.Bind(new(controller.PengajuanUjianTaController), new(*controller.PengajuanUjianTaControllerImpl)),
)

func InitializedServer() *http.Server {
	wire.Build(
		app.GetConnection,
		validator.New,
		registerSet,
		bimbinganSet,
		semproSet,
		dospemset,
		seminarhasilset,
		ujiantaset,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)
	return nil
}
