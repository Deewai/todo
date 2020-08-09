package app

import (
	"net/http"
	"encoding/json"
	"context"
	"github.com/gorilla/mux"
	pb "github.com/Deewai/todo/todo-client/internal/proto/todo"
)

type App struct{
	Service pb.TodoServiceClient
	Router *mux.Router
}

type Error struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

func NewApp(service pb.TodoServiceClient) *App{
	app := &App{Service: service, Router: mux.NewRouter().StrictSlash(true)}
	app.setRoutes()
	return app
}

func (app *App) setRoutes(){
	app.Router.HandleFunc("/todos", app.handleAddTodo).Methods("POST")
	app.Router.HandleFunc("/todos", app.handleGetTodos).Methods("GET")
	app.Router.HandleFunc("/todos/{id}", app.handleDeleteTodo).Methods("DELETE")
}

func (app *App) handleAddTodo(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	todo := pb.Todo{}
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		respondWithError(w, Error{Code: http.StatusBadRequest, Error: err.Error()})
		return
	}
	resp, err := app.Service.AddTodo(context.Background(), &todo)
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

func (app *App) handleGetTodos(w http.ResponseWriter, r *http.Request) {
	resp, err := app.Service.GetTodos(context.Background(), &pb.GetRequest{})
	if err != nil {
		respondWithError(w, Error{Code: http.StatusInternalServerError, Error: err.Error()})
		return
	}
	data := resp.GetTodos()
	if data == nil{
		respondWithJSON(w, http.StatusOK, []*pb.Todo{})
		return
	}
	respondWithJSON(w, http.StatusOK, data)
}

func (app *App) handleDeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	todoID := params["id"]
	resp, err := app.Service.DeleteTodo(context.Background(), &pb.DeleteRequest{Id: todoID})
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