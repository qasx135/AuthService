package config

import (
	"authService/internal/transport/grpc"
	"authService/pkg/logger"
	"authService/pkg/postgres"
	"context"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"os"
	"time"
)

type Config struct {
	Env            string             `yaml:"env" env-default:"local"`
	StoragePath    string             `yaml:"storage_path" env-required:"true"`
	TokenTTL       time.Duration      `yaml:"token_ttl" env-required:"true"`
	PostgresConfig postgres.Config    `yaml:"postgres" env-required:"true"`
	GRPCConfig     grpchandler.Config `yaml:"grpc" env-required:"true"`
}

func NewConfig(ctx context.Context) *Config {
	err := godotenv.Load(".env")
	if err != nil {

	}
	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		logger.GetLoggerFromCtx(ctx).Fatal(ctx, "No config path set, using default config")
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		logger.GetLoggerFromCtx(ctx).Fatal(ctx, "Config file does not exist, using default config")
	}
	var config Config
	err = cleanenv.ReadConfig(path, &config)
	if err != nil {
		logger.GetLoggerFromCtx(ctx).Fatal(ctx, "Error reading config, using default config", zap.Error(err))
		return nil
	}

	return &config
}
