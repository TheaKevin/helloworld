package todos

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTodoService(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)

	req := DataRequest{
		Task: "task 1",
	}

	todo, status, err := service.CreateTodos(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, status)
	assert.NotNil(t, todo)

}

func TestGetTodoService(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)

	req := DataRequest{
		Task: "task 1",
	}

	service.CreateTodos(req)

	todos, status, err := service.GetTodos()
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, status)
	assert.Equal(t, 1, len(todos))
	assert.Equal(t, false, todos[0].Done)
	assert.Equal(t, req.Task, todos[0].Task)

	req = DataRequest{
		Task: "task 2",
	}

	service.CreateTodos(req)

	todos, status, err = service.GetTodos()
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, status)
	assert.Equal(t, 2, len(todos))
	assert.Equal(t, false, todos[1].Done)
	assert.Equal(t, req.Task, todos[1].Task)

}

func TestDeleteTodoService(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)

	req := DataRequest{
		Task: "task 1",
	}

	service.CreateTodos(req)
	service.CreateTodos(req)

	todos, status, err := service.DeleteTodos("1")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, status)
	assert.Equal(t, 0, len(todos))

}

func TestChangeTodoService(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)

	req := DataRequest{
		Task: "task 1",
	}

	service.CreateTodos(req)
	service.CreateTodos(req)
	service.ChangeDoneTodo("1")

	todos, status, err := service.GetTodos()
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, status)
	assert.Equal(t, 2, len(todos))
	assert.Equal(t, todos[0].Done, true)
	assert.Equal(t, todos[1].Done, false)

}
