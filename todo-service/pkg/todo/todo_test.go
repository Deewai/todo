package todo

import (
	"context"
	"errors"
	"testing"

	"github.com/Deewai/todo/todo-service/internal/proto/todo"
	"github.com/stretchr/testify/assert"
)

type MyRepo struct{
	err bool
}

func (r *MyRepo) AddItem(pk string, item interface{}) error{
	if r.err{
		return errors.New("A sample error")
	}
	return nil
}

func (r *MyRepo) DeleteItem(pk string) error{
	if r.err{
		return errors.New("A sample error")
	}
	return nil
}

func (r *MyRepo) GetItems() ([]interface{}, error){
	if r.err{
		return nil, errors.New("A sample error")
	}
	return []interface{}{}, nil
}
func TestNewService(t *testing.T){
	repo := &MyRepo{}
	expected := &Service{repo}
	assert.EqualValues(t, expected, NewService(repo))
}

func TestAddTodoSuccess(t *testing.T){
	repo := &MyRepo{}
	service := &Service{repo}
	req := todo.Todo{}
	got, err := service.AddTodo(context.Background(), &req)
	assert.NoError(t, err)
	assert.True(t, got.GetCreated())
	assert.Equal(t, &req, got.GetTodo())
}

func TestAddTodoError(t *testing.T){
	repo := &MyRepo{err: true}
	service := &Service{repo}
	req := todo.Todo{}
	got, err := service.AddTodo(context.Background(), &req)
	assert.Error(t, err)
	assert.Nil(t, got)
}

func TestGetTodosSuccess(t *testing.T){
	repo := &MyRepo{}
	service := &Service{repo}
	req := todo.GetRequest{}
	got, err := service.GetTodos(context.Background(), &req)
	assert.NoError(t, err)
	assert.IsType(t, []*todo.Todo{}, got.GetTodos())
}

func TestGetTodosError(t *testing.T){
	repo := &MyRepo{err: true}
	service := &Service{repo}
	req := todo.GetRequest{}
	got, err := service.GetTodos(context.Background(), &req)
	assert.Error(t, err)
	assert.Equal(t, 0, len(got.GetTodos()))
}

func TestDeleteTodoSuccess(t *testing.T){
	repo := &MyRepo{}
	service := &Service{repo}
	req := todo.DeleteRequest{Id:""}
	got, err := service.DeleteTodo(context.Background(), &req)
	assert.NoError(t, err)
	assert.True(t, got.GetDeleted())
}

func TestDeleteTodoError(t *testing.T){
	repo := &MyRepo{err: true}
	service := &Service{repo}
	req := todo.DeleteRequest{Id:""}
	got, err := service.DeleteTodo(context.Background(), &req)
	assert.Error(t, err)
	assert.False(t, got.GetDeleted())
	assert.Nil(t, got.GetTodo())
}