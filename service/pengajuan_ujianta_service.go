package service

import (
	"context"
	"goelster/model/web"
)

type PengajuanUjianTaService interface {
	PengajuanUjianTa(ctx context.Context, request web.UjianTaRequest) web.UjianTaResponse
	FindById(ctx context.Context, dataujiantaId string) []web.UjianTaResponse
}
