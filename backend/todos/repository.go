package todos

import (
	models "github.com/TheaKevin/helloworld/model"
	"gorm.io/gorm"
)

type Repository interface {
	GetTodos() ([]models.Todos, error)
	CreateTodos(task string) (models.Todos, error)
	DeleteTodos(ID string) ([]models.Todos, error)
	ChangeDoneTodo(ID string) ([]models.Todos, error)
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

func (r *repository) CreateTodos(task string) (models.Todos, error) {
	todo := models.Todos{
		Task: task,
		Done: false,
	}

	res := r.db.Create(&todo)
	if res.Error != nil {
		return models.Todos{}, res.Error
	}

	return todo, nil
}

func (r *repository) DeleteTodos(ID string) ([]models.Todos, error) {
	var todos []models.Todos
	res := r.db.Where("ID = ?", ID).Delete(&todos)
	if res.Error != nil {
		return nil, res.Error
	}

	return todos, nil
}

func (r *repository) ChangeDoneTodo(ID string) ([]models.Todos, error) {
	var todos []models.Todos

	res := r.db.Where("ID = ?", ID).Updates(models.Todos{
		Done: true,
	})
	if res.Error != nil {
		return nil, res.Error
	}

	return todos, nil
}
