package repository

import (
	"context"
	"database/sql"
	"goelster/model/domain"
)

type PengajuanSeminarHasilRepository interface {
	UploadFile(ctx context.Context, tx *sql.Tx, hasil domain.SeminarHasil) domain.SeminarHasil
	GetFile(ctx context.Context, tx *sql.Tx) ([]domain.SeminarHasil, error)
	PengajuanJadwalSemhas(ctx context.Context, tx *sql.Tx, hasil domain.SeminarHasil) domain.SeminarHasil
	UpdateAjuanSemhas(ctx context.Context, tx *sql.Tx, hasil domain.SeminarHasil) domain.SeminarHasil
	FindById(ctx context.Context, tx *sql.Tx, id_mhs string) ([]domain.SeminarHasil, error)
}
