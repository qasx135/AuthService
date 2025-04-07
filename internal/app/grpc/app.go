package grpcapp

import (
	grpchandler "authService/internal/transport/grpc"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type App struct {
	gRPCServer *grpc.Server
	port       int
}

func NewApp(port int) *App {
	gRPCServer := grpc.NewServer()
	grpchandler.Register(gRPCServer)
	return &App{gRPCServer, port}
}

func (app *App) Run() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", app.port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	if err := app.gRPCServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func (app *App) Shutdown() {
	app.gRPCServer.GracefulStop()
}
