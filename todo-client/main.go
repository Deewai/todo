package main

import (
	"context"
	"encoding/json"
	"net/http"

	pb "github.com/Deewai/todo/todo-client/internal/proto/todo"
	"github.com/Deewai/todo/todo-client/pkg/config"
	"github.com/Deewai/todo/todo-client/pkg/log"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

type Error struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

var client pb.ShippingServiceClient

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
	client = pb.NewShippingServiceClient(conn)
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/todos", handleCreateTodo).Methods("POST")
	router.HandleFunc("/todos", handleGetTodos).Methods("GET")
	router.HandleFunc("/todos/{id}", handleDeleteTodo).Methods("DELETE")
	log.Infof("http server started on %s", ":"+conf.ClientPort)
	err = http.ListenAndServe(":"+conf.ClientPort, router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleCreateTodo(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	todo := pb.Todo{}
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		respondWithError(w, Error{Code: http.StatusBadRequest, Error: err.Error()})
		return
	}
	resp, err := client.AddTodo(context.Background(), &todo)
	if err != nil {
		respondWithError(w, Error{Code: http.StatusInternalServerError, Error: err.Error()})
		return
	}
	if !resp.GetCreated() {
		respondWithError(w, Error{Code: http.StatusInternalServerError, Error: "Error creating todo"})
		return
	}
	respondWithJSON(w, http.StatusCreated, resp.GetTodo())
}

func handleGetTodos(w http.ResponseWriter, r *http.Request) {
	resp, err := client.GetTodos(context.Background(), &pb.GetRequest{})
	if err != nil {
		respondWithError(w, Error{Code: http.StatusInternalServerError, Error: err.Error()})
		return
	}
	respondWithJSON(w, http.StatusOK, resp.GetTodos())
}

func handleDeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	todoID := params["id"]
	resp, err := client.DeleteTodo(context.Background(), &pb.DeleteRequest{Id: todoID})
	if err != nil {
		respondWithError(w, Error{Code: http.StatusInternalServerError, Error: err.Error()})
		return
	}
	if !resp.GetDeleted() {
		respondWithError(w, Error{Code: http.StatusInternalServerError, Error: "Error deleting todo"})
		return
	}
	respondWithJSON(w, http.StatusOK, resp.GetTodo())
}

func respondWithError(w http.ResponseWriter, e Error) {
	respondWithJSON(w, e.Code, e)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
