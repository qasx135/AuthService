package userservice

import (
	"authService/internal/models/user-model"
	appservice "authService/internal/services/app-service"
	"authService/internal/services/user-service/userErrors"
	"authService/pkg/jwt"
	"context"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type UserRepository interface {
	Login(ctx context.Context, email string, pass string) (token string, err error)
	GetUserFromEmail(ctx context.Context, email string) (user *user_model.User, err error)
	UserSaver(ctx context.Context, email string, pass []byte) (userID int64, err error)
	IsAdmin(ctx context.Context, user *user_model.User) (admin bool, err error)
}

type UserService struct {
	repository UserRepository
	appService *appservice.AppService
	tokenTTL   time.Duration
}

func NewUserService(repository UserRepository, service *appservice.AppService, tokenTTl time.Duration) *UserService {
	return &UserService{
		repository: repository,
		appService: service,
		tokenTTL:   tokenTTl,
	}
}
func (u *UserService) Login(ctx context.Context, email string, pass string, appID int) (token string, err error) {
	user, err := u.repository.GetUserFromEmail(ctx, email)
	if err != nil {
		if errors.Is(err, userErrors.ErrUserNotFound) {
			return "", fmt.Errorf("%w", ErrInvalidCredentials)
		}
		return "", fmt.Errorf("%w", err)
	}
	if err := bcrypt.CompareHashAndPassword(user.PassHash, []byte(pass)); err != nil {
		return "", fmt.Errorf("%w", ErrInvalidCredentials)
	}
	app, err := u.appService.GetAppFromID(ctx, appID)
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	token, err = jwt.NewToken(user, app, u.tokenTTL)
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}
	return token, nil
}
func (u *UserService) Register(ctx context.Context, email string, pass string) (userID int64, err error) {
	passHash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
	id, err := u.repository.UserSaver(ctx, email, passHash)
	if err != nil {
		return 0, err
	}
	return id, nil
}
func (u *UserService) IsAdmin(ctx context.Context, user *user_model.User) (admin bool, err error) {
	return u.repository.IsAdmin(ctx, user)
}
