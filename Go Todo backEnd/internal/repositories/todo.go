package repositories

import (
	"studioj/boilerplate_go/internal/models"

	"gorm.io/gorm"
)

type TodoRepository interface {
	FindByID(userID string, id uint) (*models.Todo, error)
	List(userID string) ([]*models.Todo, error)
	Create(userID string, todo *models.Todo) (*models.Todo, error)
	Update(userID string, todo *models.Todo) (*models.Todo, error)
	Delete(userID string, id uint) error
}

// todoRepository 구조체
type todoRepository struct {
	DB *gorm.DB
}

// NewTodoRepository - TodoRepository의 새 인스턴스 생성
func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{DB: db}
}

func (r *todoRepository) FindByID(userID string, id uint) (*models.Todo, error) {
	var todo models.Todo
	err := r.DB.Where("id = ? AND user_id = ?", id, userID).First(&todo).Error
	return &todo, err
}

func (r *todoRepository) List(userID string) ([]*models.Todo, error) {
	var todos []*models.Todo
	err := r.DB.Where("user_id = ?", userID).Find(&todos).Error
	return todos, err
}

func (r *todoRepository) Create(userID string, todo *models.Todo) (*models.Todo, error) {
	todo.UserID = userID
	err := r.DB.Create(&todo).Error
	return todo, err
}

func (r *todoRepository) Update(userID string, todo *models.Todo) (*models.Todo, error) {
	var row models.Todo
	if err := r.DB.Where("id = ? AND user_id = ?", todo.ID, userID).First(&row).Error; err != nil {
		return nil, err
	}

	err := r.DB.Model(&row).Select("*").Updates(todo).Error
	return &row, err
}

func (r *todoRepository) Delete(userID string, id uint) error {
	var todo models.Todo
	if err := r.DB.Where("id = ? AND user_id = ?", id, userID).First(&todo).Error; err != nil {
		return err
	}
	err := r.DB.Delete(&todo).Error
	return err
}
