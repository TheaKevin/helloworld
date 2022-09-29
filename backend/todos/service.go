package todos

import (
	"net/http"

	models "github.com/TheaKevin/helloworld/model"
)

type Service interface {
	GetTodos() ([]models.Todos, int, error)
	CreateTodos(req DataRequest) (models.Todos, int, error)
	DeleteTodos(ID string) ([]models.Todos, int, error)
	ChangeDoneTodo(ID string) ([]models.Todos, int, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) GetTodos() ([]models.Todos, int, error) {

	todos, err := s.repo.GetTodos()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return todos, http.StatusOK, nil
}

func (s *service) CreateTodos(req DataRequest) (models.Todos, int, error) {

	todo, err := s.repo.CreateTodos(req.Task)
	if err != nil {
		return models.Todos{}, http.StatusInternalServerError, err
	}

	return todo, http.StatusOK, nil
}

func (s *service) DeleteTodos(ID string) ([]models.Todos, int, error) {

	todos, err := s.repo.DeleteTodos(ID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return todos, http.StatusOK, nil
}

func (s *service) ChangeDoneTodo(ID string) ([]models.Todos, int, error) {

	todos, err := s.repo.ChangeDoneTodo(ID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return todos, http.StatusOK, nil
}
