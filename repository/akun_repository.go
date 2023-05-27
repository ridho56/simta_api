package repository

import (
	"context"
	"database/sql"
	"goelster/model/domain"
)

type AkunRepository interface {
	//Register(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Login(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error)
	//EmailCheck(ctx context.Context, tx *sql.Tx, email string) (string, error)
}
