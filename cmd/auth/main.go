package main

import (
	"authService/internal/app"
	"authService/internal/config"
	"authService/pkg/logger"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.Background()
	ctx, _ = logger.New(ctx)

	cfg := config.NewConfig(ctx)
	fmt.Println(cfg)
	application := app.NewApp(cfg.GRPCConfig.Port, cfg.StoragePath, cfg.TokenTTL)
	go application.GRPCSrv.Run()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	application.GRPCSrv.Shutdown()
}
