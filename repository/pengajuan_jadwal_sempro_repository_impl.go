package repository

import (
	"context"
	"database/sql"
	"errors"
	"goelster/helper"
	"goelster/model/domain"
)

type PengajuanSemproRepositoryImpl struct{}

func NewPengajuanSemproRepository() *PengajuanSemproRepositoryImpl {
	return &PengajuanSemproRepositoryImpl{}
}

func (repository *PengajuanSemproRepositoryImpl) PengajuanSempro(ctx context.Context, tx *sql.Tx, sempro domain.Sempro) domain.Sempro {
	query := "INSERT INTO simta_ujianproposal (id_ujianproposal, id_mhs, id_staf, nama_judul, abstrak, tanggal, ruang_sempro, proposalawal,status_ajuan, created_at) VALUES(?,?,?,?, ?, ?, ?, ?, ?,?);"

	res, err := tx.ExecContext(ctx, query, sempro.Id_sempro, sempro.Id_mhs, sempro.Id_staf, sempro.Judul_ta, sempro.Abstract, sempro.Tanggal, sempro.Ruang_sempro, sempro.Proposal, sempro.Status_ajuan, sempro.Created_at)

	helper.PanicIfError(err)

	id, err := res.LastInsertId()
	helper.PanicIfError(err)

	sempro.Id_sempro = string(id)
	return sempro
}

func (repository *PengajuanSemproRepositoryImpl) UpdateAjuanSempro(ctx context.Context, tx *sql.Tx, sempro domain.Sempro) domain.Sempro {
	query := "UPDATE simta_ujianproposal SET tanggal=?, ruang_sempro=? WHERE id_ujianproposal=? AND id_mhs=?;"
	_, err := tx.ExecContext(ctx, query, sempro.Tanggal, sempro.Ruang_sempro, sempro.Id_sempro, sempro.Id_mhs)
	helper.PanicIfError(err)
	return sempro
}

func (repository *PengajuanSemproRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id_mhs string) ([]domain.Sempro, error) {
	query := "SELECT id_ujianproposal, id_mhs, id_staf, nama_judul, abstrak, tanggal, ruang_sempro, status_ajuan FROM simta_ujianproposal WHERE id_mhs = ?;"
	rows, err := tx.QueryContext(ctx, query, id_mhs)
	helper.PanicIfError(err)
	defer rows.Close()

	sempros := make([]domain.Sempro, 0)
	for rows.Next() {
		sempro := domain.Sempro{}
		err := rows.Scan(&sempro.Id_sempro, &sempro.Id_mhs, &sempro.Id_staf, &sempro.Judul_ta, &sempro.Abstract, &sempro.Tanggal, &sempro.Ruang_sempro, &sempro.Status_ajuan)
		helper.PanicIfError(err)
		sempros = append(sempros, sempro)
	}

	if len(sempros) == 0 {
		return nil, errors.New("data not found")
	}
	return sempros, nil
}
