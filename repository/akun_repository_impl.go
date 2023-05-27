package repository

import (
	"context"
	"database/sql"
	"errors"
	"goelster/helper"
	"goelster/model/domain"
)

type AkunRepositoryImpl struct {
}

func NewAkunRepositoryImpl() *AkunRepositoryImpl {
	return &AkunRepositoryImpl{}
}

//func (repository *AkunRepositoryImpl) Register(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
//	script := "INSERT INTO user( nama, email, password, `level`) VALUES (?,?,?,?) "
//	result, err := tx.ExecContext(ctx, script, user.Nama, user.Email, user.Password, user.Level)
//	helper.PanicIfError(err)
//
//	id, err := result.LastInsertId()
//	helper.PanicIfError(err)
//
//	user.Id = int(id)
//	return user
//}

func (repository *AkunRepositoryImpl) Login(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {
	script := "SELECT users.id, users.email, users.username, users.password_hash, mahasiswa.id_mhs, simta_pengajuanjudul.id_staf FROM users JOIN mahasiswa ON mahasiswa.id_user = users.id JOIN simta_pengajuanjudul ON simta_pengajuanjudul.id_mhs = mahasiswa.id_mhs WHERE users.email = ? OR users.username = ? limit 1"

	rows, err := tx.QueryContext(ctx, script, user.Email, user.Username)
	helper.PanicIfError(err)
	user = domain.User{}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Email, &user.Username, &user.Password, &user.Id_mhs, &user.Id_staf)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("Email atau Password Salah")
	}
}

//func (repository *AkunRepositoryImpl) EmailCheck(ctx context.Context, tx *sql.Tx, email string) (string, error) {
//	script := "SELECT email FROM `user` WHERE email = ?"
//	rows, err := tx.QueryContext(ctx, script, email)
//	helper.PanicIfError(err)
//	defer rows.Close()
//
//	if rows.Next() {
//		return "", errors.New("email sudah digunakan")
//	} else {
//		return "email tersedia", nil
//	}
//}
