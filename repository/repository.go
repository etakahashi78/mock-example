package repository

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/etakahashi78/mock-example/model"
)

type Repository interface {
	FindAll(context.Context) ([]model.User, error)
}

type MySQLRepository struct {
	db *sqlx.DB
}

func NewMySQLRepository(db *sqlx.DB) Repository {
	return &MySQLRepository{db: db}
}

func (r MySQLRepository) FindAll(ctx context.Context) (users []model.User, err error) {
	query := `SELECT id, name, email FROM users ORDER BY id DESC`

	err = r.db.SelectContext(ctx, &users, query)
	if err != nil {
		err = fmt.Errorf("SelectContext failed. err=%w", err)
		return
	}

	return
}
