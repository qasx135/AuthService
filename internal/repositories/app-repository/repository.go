package apprepository

import (
	"authService/internal/models/app-model"
	"context"
	"github.com/jackc/pgx/v5"
)

type Repository struct {
	db *pgx.Conn
}

func (r *Repository) GetAppFromID(ctx context.Context, appID int) (*app_model.App, error) {
	panic("implement me")
}

func NewRepository(db *pgx.Conn) *Repository {
	return &Repository{db: db}
}
