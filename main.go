package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"

	"learning/grpc/configs"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	// load config
	conf := configs.Config{}
	conf.LoadConfig()

	conf.InitConnection()
	globalCtx, glbCtxCancel := context.WithCancel(context.Background())

	grpcServer, err := initGRPCServer(globalCtx, &conf)
	if err != nil {
		panic("failed to init grpc server")
	}

	// start grpc server
	go func() {
		listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", conf.App.Host, conf.App.GRPCPort))
		if err != nil {
			fmt.Println(err)
			panic("failed")
		}

		fmt.Println("GRPC server is running and listening on :", conf.App.GRPCPort)
		if err := grpcServer.Serve(listener); err != nil && !errors.Is(err, grpc.ErrServerStopped) {
			fmt.Println(err)
			panic("failed to start GRPC server")
		}
	}()

	// Keep the application running until signals trapped
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("%s signal trapped. Stopping application", <-sigChan)

	// Graceful shutdown the application
	// Notify all other background services/workers to stop
	glbCtxCancel()
	// terminate GRPC server
	grpcServer.GracefulStop()
	grpcServer.Stop()

	// Stop background jobs here
	conf.Teardown()
}

func initGRPCServer(ctx context.Context, conf *configs.Config) (*grpc.Server, error) {
	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			// recovery interceptor
			recovery.UnaryServerInterceptor(recovery.WithRecoveryHandler(func(p any) (err error) {
				fmt.Println(ctx, "recovered from panic %v, stack: %v", p, string(debug.Stack()))
				return status.Errorf(codes.Internal, "%s", p)
			})),
		),
		grpc.ChainStreamInterceptor(),
	)

	conf.InitGRPCService(server)

	return server, nil
}