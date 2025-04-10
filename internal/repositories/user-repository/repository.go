package userrepository

import (
	"authService/internal/models/user-model"
	userservice "authService/internal/services/user-service"
	"context"
	"github.com/jackc/pgx/v5"
)

type Repository struct {
	db *pgx.Conn
}

func (r *Repository) Login(ctx context.Context, email string, pass string) (string, error) {
	panic("implement me")
}

func (r *Repository) UserSaver(ctx context.Context, email string, pass []byte) (int64, error) {
	panic("implement me")
}

func (r *Repository) IsAdmin(ctx context.Context, user *user_model.User) (bool, error) {
	panic("implement me")
}

func (r *Repository) GetUserFromEmail(ctx context.Context, email string) (*user_model.User, error) {
	panic("implement me")
}

func NewRepository(db *pgx.Conn) userservice.UserRepository {
	return &Repository{
		db: db,
	}
}
