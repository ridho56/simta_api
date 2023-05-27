package repository

import (
	"context"
	"database/sql"
	"errors"
	"goelster/helper"
	"goelster/model/domain"
)

type PengajuanSeminarHasilRepositoryImpl struct{}

func NewPengajuanSeminarHasilRepository() *PengajuanSeminarHasilRepositoryImpl {
	return &PengajuanSeminarHasilRepositoryImpl{}
}

func (repository *PengajuanSeminarHasilRepositoryImpl) UploadFile(ctx context.Context, tx *sql.Tx, hasil domain.SeminarHasil) domain.SeminarHasil {
	query := "INSERT INTO simta_berkas (id_berkas, nama_berkas, file_berkas, created_at) VALUES(?, ?, ?, ?);"

	res, err := tx.ExecContext(ctx, query, hasil.Id_seminarhasil, hasil.Link_file, hasil.Created_at)

	helper.PanicIfError(err)

	id, err := res.LastInsertId()
	helper.PanicIfError(err)

	hasil.Id_seminarhasil = string(id)
	return hasil
}

func (repository *PengajuanSeminarHasilRepositoryImpl) GetFile(ctx context.Context, tx *sql.Tx) ([]domain.SeminarHasil, error) {
	query := "SELECT id_berkas, nama_berkas, file_berkas FROM simta_berkas;"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []domain.SeminarHasil
	for rows.Next() {
		var result domain.SeminarHasil
		err := rows.Scan(&result.Id_seminarhasil, &result.Link_file)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (repository *PengajuanSeminarHasilRepositoryImpl) PengajuanJadwalSemhas(ctx context.Context, tx *sql.Tx, hasil domain.SeminarHasil) domain.SeminarHasil {
	query := "INSERT INTO simta_seminarhasil (id_seminarhasil, id_mhs, id_staf,nama_judul, abstrak, jadwal_semhas,ruang_semhas, link_file, status_ajuan) VALUES(?,?,?, ?, ?, ?, ?, ?, ?);"

	res, err := tx.ExecContext(ctx, query, hasil.Id_seminarhasil, hasil.Id_mhs, hasil.Id_staf, hasil.Nama_judul, hasil.Abstak, hasil.Jadwal_semhas, hasil.Ruang_semhas, hasil.Link_file, hasil.Status_ajuan)

	helper.PanicIfError(err)

	id, err := res.LastInsertId()
	helper.PanicIfError(err)

	hasil.Id_seminarhasil = string(id)
	return hasil
}

func (repository *PengajuanSeminarHasilRepositoryImpl) UpdateAjuanSemhas(ctx context.Context, tx *sql.Tx, hasil domain.SeminarHasil) domain.SeminarHasil {
	query := "UPDATE simta_seminarhasil SET jadwal_semhas=?, ruang_semhas=? WHERE id_seminarhasil=?;"
	_, err := tx.ExecContext(ctx, query, hasil.Jadwal_semhas, hasil.Ruang_semhas, hasil.Id_seminarhasil)
	helper.PanicIfError(err)
	return hasil
}

func (repository *PengajuanSeminarHasilRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id_mhs string) ([]domain.SeminarHasil, error) {
	query := "SELECT id_seminarhasil, id_mhs, id_staf, nama_judul, abstrak, jadwal_semhas, ruang_semhas, status_ajuan FROM simta_seminarhasil WHERE id_mhs = ?;"
	rows, err := tx.QueryContext(ctx, query, id_mhs)
	helper.PanicIfError(err)
	defer rows.Close()

	semhass := make([]domain.SeminarHasil, 0)
	for rows.Next() {
		semhas := domain.SeminarHasil{}
		err := rows.Scan(&semhas.Id_seminarhasil, &semhas.Id_mhs, &semhas.Id_staf, &semhas.Nama_judul, &semhas.Abstak, &semhas.Jadwal_semhas, &semhas.Ruang_semhas, &semhas.Status_ajuan)
		helper.PanicIfError(err)
		semhass = append(semhass, semhas)
	}

	if len(semhass) == 0 {
		return nil, errors.New("data not found")
	}
	return semhass, nil
}
