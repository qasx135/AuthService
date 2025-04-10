package repositories

import (
	apprepository "authService/internal/repositories/app-repository"
	userrepository "authService/internal/repositories/user-repository"
	appservice "authService/internal/services/app-service"
	userservice "authService/internal/services/user-service"
	"github.com/jackc/pgx/v5"
)

type Repositories struct {
	UserRepository userservice.UserRepository
	AppRepository  appservice.AppRepository
}

func (r *Repositories) Register(db *pgx.Conn) *Repositories {
	r.UserRepository = userrepository.NewRepository(db)
	r.AppRepository = apprepository.NewRepository(db)
	return &Repositories{}
}
