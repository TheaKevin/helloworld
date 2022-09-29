package todos

import (
	models "github.com/TheaKevin/helloworld/model"
	"gorm.io/gorm"
)

type Repository interface {
	GetTodos() ([]models.Todos, error)
	CreateTodos(data DataRequest) (models.Todos, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetTodos() ([]models.Todos, error) {
	var todos []models.Todos
	res := r.db.Find(&todos)
	if res.Error != nil {
		return nil, res.Error
	}

	return todos, nil
}

func (r *repository) CreateTodos(req DataRequest) (models.Todos, error) {
	todo := models.Todos{
		Task: req.Task,
		Done: false,
	}

	res := r.db.Create(&todo)
	if res.Error != nil {
		return models.Todos{}, res.Error
	}

	return todo, nil
}
