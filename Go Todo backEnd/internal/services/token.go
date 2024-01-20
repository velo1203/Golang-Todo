package services

import (
	"fmt"
	"strings"

	config "studioj/boilerplate_go/configs"
	"studioj/boilerplate_go/internal/models"
	"studioj/boilerplate_go/internal/repositories"

	"github.com/golang-jwt/jwt/v5"
)

type tokenService struct {
	repository repositories.TokenRepository
}

type TokenService interface {
	Find(string) (*models.Token, error)
	Create(*models.Token) (*models.Token, error)
	Update(*models.Token) (*models.Token, error)
	Delete(string) error

	Generate(string, int64) (string, error)
	Verify(tokenString string) (*jwt.Token, error)

	GetJwtToken(string) (*jwt.Token, error)
	ExtractTokenString(string) string
	VerifyTokenString(string) (*jwt.Token, error)
	ValidToken(*jwt.Token) (bool, error)
}

func NewTokenService(repository repositories.TokenRepository) TokenService {
	return &tokenService{
		repository: repository,
	}
}

func (s *tokenService) Find(id string) (*models.Token, error) {
	return s.repository.Find(id)
}

func (s *tokenService) Create(token *models.Token) (*models.Token, error) {
	return s.repository.Create(token)
}

func (s *tokenService) Update(token *models.Token) (*models.Token, error) {
	return s.repository.Update(token)
}

func (s *tokenService) Delete(id string) error {
	return s.repository.Delete(id)
}

func (s *tokenService) Verify(tokenString string) (*jwt.Token, error) {
	jwtToken, err := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", jwtToken.Header["alg"])
		}
		return []byte(config.SECRET_KEY), nil
	})

	return jwtToken, err
}

func (s *tokenService) Generate(user_id string, expire_at int64) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["sub"] = user_id
	claims["exp"] = expire_at
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := at.SignedString([]byte(config.SECRET_KEY))

	return token, err
}

func (s *tokenService) GetJwtToken(tokenString string) (*jwt.Token, error) {
	jwtToken, err := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", jwtToken.Header["alg"])
		}
		return []byte(config.SECRET_KEY), nil
	})
	return jwtToken, err
}

func (s *tokenService) ExtractTokenString(authorization string) string {
	strArr := strings.Split(authorization, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func (s *tokenService) VerifyTokenString(tokenString string) (*jwt.Token, error) {
	jwtToken, err := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", jwtToken.Header["alg"])
		}
		return []byte(config.SECRET_KEY), nil
	})

	if err != nil {
		return nil, err
	}

	return jwtToken, nil
}

func (s *tokenService) ValidToken(jwtToken *jwt.Token) (bool, error) {
	if jwtToken == nil {
		return false, fmt.Errorf("no token")
	}
	return jwtToken.Valid, nil
}
