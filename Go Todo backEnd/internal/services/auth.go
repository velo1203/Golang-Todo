package services

import (
	"errors"
	"studioj/boilerplate_go/internal/models"
	"studioj/boilerplate_go/internal/repositories"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	userRepository  repositories.UserRepository
	tokenRepository repositories.TokenRepository
}

type AuthService interface {
	Register(*models.RegisterRequest) (*models.User, error)
	Login(*models.LoginRequest) (*models.User, error)
	CreateToken(string) (*models.Token, error)
}

func NewAuthService(userRepository repositories.UserRepository, tokenRepository repositories.TokenRepository) AuthService {
	return &authService{
		userRepository:  userRepository,
		tokenRepository: tokenRepository,
	}
}

func (s *authService) Register(request *models.RegisterRequest) (*models.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	if err != nil {
		return nil, errors.New("fail to hash password")
	}

	// Create the user
	newUser := models.User{
		ID:       uuid.NewString(),
		Username: request.Username,
		Password: string(hash),
	}

	user, err := s.userRepository.Create(&newUser)
	if err != nil {
		return nil, errors.New("fail to create user")
	}

	return user, err
}

func (s *authService) Login(request *models.LoginRequest) (*models.User, error) {
	user, err := s.userRepository.FindByUsername(request.Username)
	if err != nil || user == nil {
		return nil, errors.New("invalid user or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return nil, errors.New("invalid user or password")
	}
	return user, nil
}

func (s *authService) CreateToken(user_id string) (*models.Token, error) {
	tokenExpiredAt := time.Now().Add(time.Hour * 24 * 30)
	accessToken, err := s.tokenRepository.Generate(user_id, tokenExpiredAt.Unix())
	if err != nil {
		return nil, err
	}

	refreshExpiredAt := time.Now().Add(time.Hour * 24 * 90)
	refreshToken, err := s.tokenRepository.Generate(user_id, refreshExpiredAt.Unix())
	if err != nil {
		return nil, err
	}

	newToken := &models.Token{
		ID:           uuid.NewString(),
		UserID:       user_id,
		Token:        accessToken,
		RefreshToken: refreshToken,
		Status:       "valid",
		ExpireAt:     tokenExpiredAt,
	}

	token, err := s.tokenRepository.Create(newToken)

	return token, err
}
