package app

import (
	"authService/internal/app/grpc"
	"time"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func NewApp(grpcPort int,
	storagePath string,
	tokenTTL time.Duration) *App {

	grpcApp := grpcapp.NewApp(grpcPort)
	return &App{GRPCSrv: grpcApp}
}
