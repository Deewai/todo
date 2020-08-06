package todo

import (
	"context"
	"github.com/Deewai/todo/proto/build/todo"
	"github.com/google/uuid"
)

// Repository interface to be satisfied by the storage module for managing/persisting todos
type Repository interface{
	AddItem(string, interface{}) error
	DeleteItem(string) error
	GetItems() ([]interface{}, error)
}

// Service satisfies the generated todo pb service
type Service struct{
	repository Repository
}

func NewService(repository Repository) *Service{
	return &Service{repository: repository}
}

	
func(s *Service) AddTodo(ctx context.Context, req *todo.Todo) (*todo.AddResponse, error){
	req.Id = uuid.New().String()
	err := s.repository.AddItem(req.Id,*req)
	if err != nil{
		return nil, err
	}
	return &todo.AddResponse{Created: true, Todo: req}, nil
}

func(s *Service) GetTodos(ctx context.Context, req *todo.GetRequest) (*todo.GetResponse, error){
	items, err := s.repository.GetItems()
	if err != nil{
		return nil, err
	}
	todos := []*todo.Todo{}
	for _, val := range items {
		todo := val.(todo.Todo)
		todos = append(todos, &todo)
	}
	return &todo.GetResponse{Todos: todos}, nil
}

func(s *Service) DeleteTodo(ctx context.Context, req *todo.DeleteRequest) (*todo.DeleteResponse, error){
	todoID := req.Id
	err := s.repository.DeleteItem(todoID)
	if err != nil{
		return nil, err
	}
	return &todo.DeleteResponse{Deleted: true}, nil
}