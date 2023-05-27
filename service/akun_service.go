package service

import (
	"context"
	"goelster/model/web"
)

type AkunService interface {
	//Register(ctx context.Context, request web.RegisterRequest) web.RegisterResponse
	Login(ctx context.Context, request web.LoginRequest) web.LoginResponse
	//EmailCheck(ctx context.Context, email string) web.EmailCheckResponse
}
