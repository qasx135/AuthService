package grpchandler

import (
	"context"
	authv1 "github.com/qasx135/protos/gen/go/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"regexp"
	"time"
)

const (
	UserEmail = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
)

type Config struct {
	Port    int           `yaml:"port" env-required:"true"`
	Timeout time.Duration `yaml:"timeout" env-required:"true"`
}

type ServerAPI struct {
	authv1.UnimplementedAuthServer
}

func Register(gRPC *grpc.Server) {
	authv1.RegisterAuthServer(gRPC, &ServerAPI{})
}

func (srv *ServerAPI) Login(ctx context.Context,
	in *authv1.LoginRequest) (*authv1.LoginResponse, error) {
	re := regexp.MustCompile(UserEmail)
	if !re.MatchString(in.GetEmail()) {
		return nil, status.Error(codes.InvalidArgument, "invalid email")
	}
	return &authv1.LoginResponse{
		Token: "",
	}, nil
}

func (srv *ServerAPI) Register(ctx context.Context,
	in *authv1.RegisterRequest) (*authv1.RegisterResponse, error) {
	panic("implement me")
}

func (srv *ServerAPI) IsAdmin(ctx context.Context,
	in *authv1.IsAdminRequest) (*authv1.IsAdminResponse, error) {
	panic("implement me")
}
