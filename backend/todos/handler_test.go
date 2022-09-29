package todos

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	models "github.com/TheaKevin/helloworld/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateTodoHandler(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/send", handler.CreateTodo)
	payload := `{"task": "task 1"}`
	req, err := http.NewRequest("POST", "/send", strings.NewReader(payload))
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	type response struct {
		Message string       `json:"message"`
		Data    models.Todos `json:"data"`
	}
	var res response
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
	assert.Equal(t, "success", res.Message)
	assert.Equal(t, "task 1", res.Data.Task)
}

func TestGetTodoHandler(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	db.Create(&models.Todos{
		Task: "task 1",
	})

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/", handler.GetTodos)
	req, err := http.NewRequest("GET", "/", nil)
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	type response struct {
		Message string         `json:"message"`
		Data    []models.Todos `json:"data"`
	}
	var res response
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
	assert.Equal(t, "success", res.Message)
	assert.Equal(t, "task 1", res.Data[0].Task)
}

func TestDeleteTodoHandler(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/send", handler.CreateTodo)
	payload := `{"task": "task 1"}`
	http.NewRequest("POST", "/send", strings.NewReader(payload))
	r.DELETE("/delete/:taskId", handler.DeleteTodo)
	req, err := http.NewRequest("DELETE", "/delete/:taskId", strings.NewReader("1"))
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	type response struct {
		Message string       `json:"message"`
		Data    models.Todos `json:"data"`
	}
	var res response
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
	assert.Equal(t, "Task Deleted", res.Message)
	assert.Equal(t, "", res.Data.Task)
}

func TestChangeDoneTodoHandler(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/send", handler.CreateTodo)
	payload := `{"task": "task 1"}`
	http.NewRequest("POST", "/send", strings.NewReader(payload))
	r.PUT("/changeDone/:taskId", handler.ChangeDoneTodo)
	req, err := http.NewRequest("PUT", "/changeDone/:taskId", strings.NewReader("1"))
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	type response struct {
		Message string       `json:"message"`
		Data    models.Todos `json:"data"`
	}
	var res response
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
	assert.Equal(t, "Task Update Success", res.Message)
}
