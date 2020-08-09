package app

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	pb "github.com/Deewai/todo/todo-client/internal/proto/todo"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

var ErrSample = errors.New("a sample error")

type myService struct {
	err bool
}

func (s *myService) AddTodo(ctx context.Context, in *pb.Todo, opts ...grpc.CallOption) (*pb.AddResponse, error) {
	if s.err {
		return nil, ErrSample
	}
	return &pb.AddResponse{Created: true, Todo: nil}, nil
}
func (s *myService) GetTodos(ctx context.Context, in *pb.GetRequest, opts ...grpc.CallOption) (*pb.GetResponse, error) {
	if s.err {
		return nil, ErrSample
	}
	return &pb.GetResponse{Todos: []*pb.Todo{&pb.Todo{Name: "a name", Description: "a description", Deadline: "a deadline"}}}, nil
}
func (s *myService) DeleteTodo(ctx context.Context, in *pb.DeleteRequest, opts ...grpc.CallOption) (*pb.DeleteResponse, error) {
	if s.err {
		return nil, ErrSample
	}
	return &pb.DeleteResponse{Deleted: true, Todo: nil}, nil
}

func TestAddTodoSuccess(t *testing.T) {
	a := App{}
	a.Service = &myService{}
	a.Router = mux.NewRouter()
	a.setRoutes()
	body := bytes.NewBuffer([]byte(`{"name":"a name","description":"a description", "deadline":"a deadline"}`))
	req, _ := http.NewRequest("POST", "/todos", body)
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusCreated, rr.Code)
}
func TestAddTodoBadRequest(t *testing.T) {
	a := App{}
	a.Service = &myService{}
	a.Router = mux.NewRouter()
	a.setRoutes()
	body := bytes.NewBuffer([]byte(`{"name":"a name",description:"a description", "deadline":"a deadline"}`))
	req, _ := http.NewRequest("POST", "/todos", body)
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestAddTodoNotCreated(t *testing.T) {
	a := App{}
	a.Service = &myService{err: true}
	a.Router = mux.NewRouter()
	a.setRoutes()
	body := bytes.NewBuffer([]byte(`{"name":"a name","description":"a description", "deadline":"a deadline"}`))
	req, _ := http.NewRequest("POST", "/todos", body)
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	var m map[string]interface{}
	json.Unmarshal(rr.Body.Bytes(), &m)
	assert.NotNil(t, m["error"])
}

func TestGetTodosSuccess(t *testing.T) {
	a := App{}
	a.Service = &myService{}
	a.Router = mux.NewRouter()
	a.setRoutes()
	req, _ := http.NewRequest("GET", "/todos", nil)
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	var m []*pb.Todo
	assert.NoError(t, json.Unmarshal(rr.Body.Bytes(), &m))
	assert.GreaterOrEqual(t, len(m), 1)
}

func TestGetTodosError(t *testing.T) {
	a := App{}
	a.Service = &myService{err: true}
	a.Router = mux.NewRouter()
	a.setRoutes()
	req, _ := http.NewRequest("GET", "/todos", nil)
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

func TestDeleteTodoSuccess(t *testing.T) {
	a := App{}
	a.Service = &myService{}
	a.Router = mux.NewRouter()
	a.setRoutes()
	// body := bytes.NewBuffer([]byte(""))
	req, _ := http.NewRequest("DELETE", "/todos/1", nil)
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestDeleteTodoError(t *testing.T) {
	a := App{}
	a.Service = &myService{err: true}
	a.Router = mux.NewRouter()
	a.setRoutes()
	req, _ := http.NewRequest("GET", "/todos", nil)
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}
