package repositories

import (
	config "studioj/boilerplate_go/configs"
	"studioj/boilerplate_go/internal/models"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type tokenRepository struct {
	DB *gorm.DB
}

func NewTokenRepository(db *gorm.DB) TokenRepository {
	return &tokenRepository{
		DB: db,
	}
}

type TokenRepository interface {
	Generate(string, int64) (string, error)
	Find(string) (*models.Token, error)
	Create(*models.Token) (*models.Token, error)
	Update(*models.Token) (*models.Token, error)
	Delete(string) error
}

func (s *tokenRepository) Generate(sub string, expire_at int64) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["sub"] = sub
	claims["exp"] = expire_at
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := at.SignedString([]byte(config.SECRET_KEY))

	return token, err
}

func (r *tokenRepository) Find(id string) (*models.Token, error) {
	var token *models.Token
	err := r.DB.Where("id = ?", id).First(&token).Error
	return token, err
}

func (r *tokenRepository) Create(token *models.Token) (*models.Token, error) {
	err := r.DB.Create(&token).Error
	return token, err
}

func (r *tokenRepository) Update(token *models.Token) (*models.Token, error) {
	var row *models.Token
	if err := r.DB.Where("id=?", token.ID).First(&row).Error; err != nil {
		return nil, err
	}

	err := r.DB.Model(&row).Select("*").Updates(&token).Error
	return row, err
}

func (r *tokenRepository) Delete(id string) error {
	var token *models.Token
	if err := r.DB.Where("id=?", id).First(&token).Error; err != nil {
		return err
	}
	err := r.DB.Delete(&token).Error
	return err
}
