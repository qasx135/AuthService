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

type AuthService interface {
	Login(ctx context.Context, email string, password string, appID int32) (token string, err error)
	Register(ctx context.Context, email string, password string) (userID int64, err error)
	IsAdmin(ctx context.Context, userID int64) (admin bool, err error)
}
type Config struct {
	Port    int           `yaml:"port" env-required:"true"`
	Timeout time.Duration `yaml:"timeout" env-required:"true"`
}

type ServerAPI struct {
	authv1.UnimplementedAuthServer
	auth AuthService
}

func Register(gRPC *grpc.Server, auth AuthService) {
	authv1.RegisterAuthServer(gRPC, &ServerAPI{auth: auth})
}

func (srv *ServerAPI) Login(ctx context.Context,
	in *authv1.LoginRequest) (*authv1.LoginResponse, error) {
	re := regexp.MustCompile(UserEmail)
	if !re.MatchString(in.GetEmail()) {
		return nil, status.Error(codes.InvalidArgument, "invalid email")
	}
	if in.GetPassword() == "" {
		return nil, status.Error(codes.InvalidArgument, "Invalid password")
	}
	if in.GetAppId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "Invalid app ID")
	}
	token, err := srv.auth.Login(ctx, in.GetEmail(), in.GetPassword(), in.GetAppId())
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to login")
	}
	return &authv1.LoginResponse{
		Token: token,
	}, nil
}

func (srv *ServerAPI) Register(ctx context.Context,
	in *authv1.RegisterRequest) (*authv1.RegisterResponse, error) {
	re := regexp.MustCompile(UserEmail)
	if !re.MatchString(in.GetEmail()) {
		return nil, status.Error(codes.InvalidArgument, "invalid email")
	}
	if in.GetPassword() == "" {
		return nil, status.Error(codes.InvalidArgument, "Invalid password")
	}
	userID, err := srv.auth.Register(ctx, in.GetEmail(), in.GetPassword())
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to register")
	}
	return &authv1.RegisterResponse{
		UserId: userID,
	}, nil
}

func (srv *ServerAPI) IsAdmin(ctx context.Context,
	in *authv1.IsAdminRequest) (*authv1.IsAdminResponse, error) {
	if in.GetUserId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "Invalid user ID")
	}
	isAdmin, err := srv.auth.IsAdmin(ctx, in.GetUserId())
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to check admin")
	}
	return &authv1.IsAdminResponse{
		IsAdmin1: isAdmin,
	}, nil
}
