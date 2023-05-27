package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"goelster/exception"
	"goelster/helper"
	"goelster/model/domain"
	"goelster/model/web"
	"goelster/repository"
)

type DospemServiceImpl struct {
	DospemRepository repository.DospemRepository
	DB               *sql.DB
	Validate         *validator.Validate
}

func NewDospemService(dospemRepository repository.DospemRepository, DB *sql.DB, validate *validator.Validate) *DospemServiceImpl {
	return &DospemServiceImpl{
		DospemRepository: dospemRepository,
		DB:               DB,
		Validate:         validate}
}

func (service *DospemServiceImpl) SetDospem(ctx context.Context, request web.DospemRequest) web.DospemResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	dospem := domain.SetDospem{
		Id_mhs:  request.Id_mhs,
		Id_staf: request.Id_staf,
	}
	dospem = service.DospemRepository.SetDospem(ctx, tx, dospem)
	return helper.ToDospemResponse(dospem)
}

func (service *DospemServiceImpl) GetDospem(ctx context.Context, dosen string) []web.DospemResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	datadosen, err := service.DospemRepository.GetDospem(ctx, tx, dosen)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToNewDospemResponse(datadosen)
}

func (service *DospemServiceImpl) GetMhs(ctx context.Context) []web.NamaIDMhs {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	hasil, err := service.DospemRepository.GetMhs(ctx, tx)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToNewMhsResponse(hasil)
}

func (service *DospemServiceImpl) GetDosen(ctx context.Context) []web.NamaIDDosen {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	result, err := service.DospemRepository.GetDosen(ctx, tx)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToNewDosenResponse(result)
}
