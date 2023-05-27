package service

import (
	"context"
	"goelster/model/web"
)

type DospemService interface {
	SetDospem(ctx context.Context, request web.DospemRequest) web.DospemResponse
	GetDospem(ctx context.Context, dosen string) []web.DospemResponse
	GetMhs(ctx context.Context) []web.NamaIDMhs
	GetDosen(ctx context.Context) []web.NamaIDDosen
}
