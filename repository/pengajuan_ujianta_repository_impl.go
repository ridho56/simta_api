package repository

import (
	"context"
	"database/sql"
	"errors"
	"goelster/helper"
	"goelster/model/domain"
)

type PengajuanUjianTaRepositoryImpl struct{}

func NewPengajuanUjianTaRepository() *PengajuanUjianTaRepositoryImpl {
	return &PengajuanUjianTaRepositoryImpl{}
}

func (repository *PengajuanUjianTaRepositoryImpl) PengajuanUjianTa(ctx context.Context, tx *sql.Tx, ujian domain.UjianTa) domain.UjianTa {
	query := "INSERT INTO simta_ujianta (id_ujianta, id_mhs, id_staf, nama_judul, abstrak, ruangan, tanggal,status, proposalakhir) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?);"

	res, err := tx.ExecContext(ctx, query, ujian.Id_ujianta, ujian.Id_mhs, ujian.Id_staf, ujian.Nama_judul, ujian.Abstrak, ujian.Ruangan, ujian.Tanggal, ujian.Status_ajuan, ujian.Proposalakhir)

	helper.PanicIfError(err)

	id, err := res.LastInsertId()
	helper.PanicIfError(err)

	ujian.Id_ujianta = string(id)
	return ujian
}

func (repository *PengajuanUjianTaRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id_mhs string) ([]domain.UjianTa, error) {
	query := "SELECT id_ujianta, id_mhs, id_staf, nama_judul, abstrak, ruangan, tanggal, status_ajuan FROM simta_ujianta WHERE id_mhs = ?;"
	rows, err := tx.QueryContext(ctx, query, id_mhs)
	helper.PanicIfError(err)
	defer rows.Close()

	ujiantaa := make([]domain.UjianTa, 0)
	for rows.Next() {
		ujianta := domain.UjianTa{}
		err := rows.Scan(&ujianta.Id_ujianta, &ujianta.Id_mhs, &ujianta.Id_staf, &ujianta.Nama_judul, &ujianta.Abstrak, &ujianta.Ruangan, &ujianta.Tanggal, &ujianta.Status_ajuan)
		helper.PanicIfError(err)
		ujiantaa = append(ujiantaa, ujianta)
	}

	if len(ujiantaa) == 0 {
		return nil, errors.New("data not found")
	}
	return ujiantaa, nil
}
