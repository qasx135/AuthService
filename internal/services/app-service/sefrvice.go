package appservice

import (
	"authService/internal/models/app-model"
	"context"
)

type AppRepository interface {
	GetAppFromID(ctx context.Context, appID int) (*app_model.App, error)
}

type AppService struct {
	appRepository AppRepository
}

func NewAppService(appRepository AppRepository) *AppService {
	return &AppService{appRepository: appRepository}
}
func (s *AppService) GetAppFromID(ctx context.Context, appID int) (*app_model.App, error) {
	panic("implement me")
	//app, err := s.appRepository.GetAppFromID(ctx, appID)
	//if err != nil {
	//	return nil, err
	//}
	//return app, nil
}
