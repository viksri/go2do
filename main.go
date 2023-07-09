package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"todo/service"

	"go.uber.org/fx"
	"google.golang.org/grpc"
)

const servicePort int = 50051

func main() {
	fx.New(
		fx.Provide(newServer),
		fx.Invoke(lifeCycle),
	).Run()
}

func newServer() *grpc.Server {
	server := grpc.NewServer()
	service.RegisterTaskProtoServer(server, &service.TaskAppServer{})
	return server
}

func lifeCycle(lc fx.Lifecycle, server *grpc.Server) {
	fmt.Println("Starting lifecycle")
	listener, err := net.Listen("tcp", "localhost:"+strconv.Itoa(servicePort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				var err error
				go func() {
					err = server.Serve(listener)
				}()
				return err
			},
			OnStop: func(ctx context.Context) error {
				server.Stop()
				return nil
			},
		},
	)
}
