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

type AkunServiceImpl struct {
	AkunRepository repository.AkunRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewAkunService(akunRepository repository.AkunRepository, DB *sql.DB, validate *validator.Validate) *AkunServiceImpl {
	return &AkunServiceImpl{
		AkunRepository: akunRepository,
		DB:             DB,
		Validate:       validate}
}

//func (service *AkunServiceImpl) Register(ctx context.Context, request web.RegisterRequest) web.RegisterResponse {
//	err := service.Validate.Struct(request)
//	helper.PanicIfError(err)
//
//	tx, err := service.DB.Begin()
//	helper.PanicIfError(err)
//	defer helper.CommitOrRollback(tx)
//
//	pwhash, err := helper.HashPassword(request.Password)
//	helper.PanicIfError(err)
//
//	register := domain.User{
//		Nama:     request.Nama,
//		Email:    request.Email,
//		Password: pwhash,
//		Level:    request.Level,
//	}
//	register = service.AkunRepository.Register(ctx, tx, register)
//	return helper.ToRegisterResponse(register)
//}

func (service *AkunServiceImpl) Login(ctx context.Context, request web.LoginRequest) web.LoginResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	email := domain.User{Email: request.EmailOrUsername, Username: request.EmailOrUsername}

	login, err := service.AkunRepository.Login(ctx, tx, email)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	result := helper.CompareHashAndPasswordSha384(login.Password, request.Password)
	if !result {
		panic(exception.NewNotFoundError("Email atau Password Salah"))
	}

	login.Email = request.EmailOrUsername

	return helper.ToLoginResponse(login)
}

//func (service *AkunServiceImpl) EmailCheck(ctx context.Context, email string) web.EmailCheckResponse {
//	tx, err := service.DB.Begin()
//	helper.PanicIfError(err)
//	defer helper.CommitOrRollback(tx)
//
//	emailCheck, err := service.AkunRepository.EmailCheck(ctx, tx, email)
//	if err != nil {
//		panic(exception.NewNotFoundError(err.Error()))
//	}
//
//	return helper.ToEmailCheckResponse(emailCheck)
//}
