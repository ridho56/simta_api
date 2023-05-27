package service

import (
	"context"
	"goelster/model/web"
)

type PengajuanSemproService interface {
	PengajuanSempro(ctx context.Context, request web.SemproRequest) web.SemproResponse
	UpdateAjuanSempro(ctx context.Context, request []web.SemproUpdateRequest) []web.SemproResponse
	FindById(ctx context.Context, dataasemproId string) []web.SemproResponse
}
