package main

import (
	"net"

	pb "github.com/Deewai/todo/todo-service/internal/proto/todo"
	"github.com/Deewai/todo/todo-service/pkg/config"
	"github.com/Deewai/todo/todo-service/pkg/log"
	"github.com/Deewai/todo/todo-service/pkg/storage"
	"github.com/Deewai/todo/todo-service/pkg/todo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	env, err := config.LoadEnv()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	db := storage.NewDB()
	service := todo.NewService(db)
	lis, err := net.Listen("tcp", ":"+env.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterShippingServiceServer(s, service)

	reflection.Register(s)
	log.Infof("Serving on port :%s", env.Port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
