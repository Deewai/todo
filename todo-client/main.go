package main

import (
	"net/http"

	pb "github.com/Deewai/todo/todo-client/internal/proto/todo"
	"github.com/Deewai/todo/todo-client/pkg/config"
	"github.com/Deewai/todo/todo-client/pkg/log"
	"github.com/Deewai/todo/todo-client/pkg/app"

	"google.golang.org/grpc"
)



func main() {
	conf, err := config.LoadEnv()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	// Set up a connection to the server.
	conn, err := grpc.Dial(conf.ServerAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewTodoServiceClient(conn)

	App := app.NewApp(client)

	log.Infof("http server started on %s", ":"+conf.ClientPort)
	err = http.ListenAndServe(":"+conf.ClientPort, App.Router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
