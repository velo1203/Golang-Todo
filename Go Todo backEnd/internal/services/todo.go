package services

import (
	"errors"
	"studioj/boilerplate_go/internal/models"
	"studioj/boilerplate_go/internal/repositories"
)

// TodoService 인터페이스 정의
type TodoService interface {
	Create(userID string, request *models.CreateRequest) (*models.Todo, error)
	ReadByID(userID string, id uint) (*models.Todo, error)
	ReadAll(userID string) ([]*models.Todo, error)
	Update(userID string, request *models.UpdateRequest) (*models.Todo, error)
	Delete(userID string, id uint) error
}

// todoService 구조체 정의
type todoService struct {
	Todorepository repositories.TodoRepository
}

// NewTodoService 함수는 TodoService 인스턴스를 생성합니다.
func NewTodoService(repository repositories.TodoRepository) TodoService {
	return &todoService{
		Todorepository: repository,
	}
}

// Create 메소드 구현
func (s *todoService) Create(userID string, request *models.CreateRequest) (*models.Todo, error) {
	newTodo := models.Todo{
		UserID:    userID,
		Title:     request.Title,
		Completed: false,
	}

	todo, err := s.Todorepository.Create(userID, &newTodo)
	if err != nil {
		return nil, errors.New("fail to create todo")
	}

	return todo, nil
}

// ReadByID 메소드 구현
func (s *todoService) ReadByID(userID string, id uint) (*models.Todo, error) {
	todo, err := s.Todorepository.FindByID(userID, id)
	if err != nil {
		return nil, err
	}
	if todo.UserID != userID {
		return nil, errors.New("unauthorized access to todo")
	}
	return todo, nil
}

// ReadAll 메소드 구현
func (s *todoService) ReadAll(userID string) ([]*models.Todo, error) {
	// userID를 사용하여 사용자별 투두 항목을 필터링합니다.
	todos, err := s.Todorepository.List(userID)
	if err != nil {
		return nil, err
	}
	return todos, nil
}

// Update 메소드 구현
func (s *todoService) Update(userID string, request *models.UpdateRequest) (*models.Todo, error) {
	existingTodo, err := s.Todorepository.FindByID(userID, request.ID)
	if err != nil {
		return nil, err
	}
	if existingTodo.UserID != userID {
		return nil, errors.New("unauthorized access to todo")
	}

	updatedTodo := models.Todo{
		ID:        request.ID,
		Title:     existingTodo.Title, // 기존 Title 값을 유지
		Completed: request.Completed,
		UserID:    userID,
	}

	todo, err := s.Todorepository.Update(userID, &updatedTodo)
	if err != nil {
		return nil, errors.New("fail to update todo")
	}

	return todo, nil
}

// Delete 메소드 구현
func (s *todoService) Delete(userID string, id uint) error {
	todo, err := s.Todorepository.FindByID(userID, id)
	if err != nil {
		return err
	}
	if todo.UserID != userID {
		return errors.New("unauthorized access to delete todo")
	}

	return s.Todorepository.Delete(userID, id)
}
