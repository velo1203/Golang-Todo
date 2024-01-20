package services

import (
	"studioj/boilerplate_go/internal/models"
	"studioj/boilerplate_go/internal/repositories"
)

type userService struct {
	repository      repositories.UserRepository
	tokenRepository repositories.TokenRepository
}

type UserService interface {
	FindByID(string) (*models.User, error)
	FindByUsername(string) (*models.User, error)
	Create(*models.User) (*models.User, error)
}

func NewUserService(repository repositories.UserRepository, tokenRepository repositories.TokenRepository) UserService {
	return &userService{
		repository:      repository,
		tokenRepository: tokenRepository,
	}
}

func (s *userService) FindByID(id string) (*models.User, error) {
	return s.repository.FindByID(id)
}

func (s *userService) FindByUsername(username string) (*models.User, error) {
	return s.repository.FindByUsername(username)
}

func (s *userService) Create(user *models.User) (*models.User, error) {
	result, err := s.repository.Create(user)

	return result, err
}
