package repositories

import (
	"studioj/boilerplate_go/internal/models"

	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		DB: db,
	}
}

type UserRepository interface {
	List() (*[]models.User, error)
	FindByID(string) (*models.User, error)
	FindByUsername(string) (*models.User, error)
	Create(*models.User) (*models.User, error)
	Update(*models.User) (*models.User, error)
	Delete(string) error
}

func (r *userRepository) List() (*[]models.User, error) {
	var users *[]models.User
	err := r.DB.Find(&users).Error
	return users, err
}

func (r *userRepository) FindByID(id string) (*models.User, error) {
	var user *models.User
	err := r.DB.Where("id = ?", id).First(&user).Error
	return user, err
}

func (r *userRepository) FindByUsername(username string) (*models.User, error) {
	var user *models.User
	err := r.DB.Where("username = ?", username).First(&user).Error
	return user, err
}

func (r *userRepository) Create(user *models.User) (*models.User, error) {
	err := r.DB.Create(&user).Error
	return user, err
}

func (r *userRepository) Update(user *models.User) (*models.User, error) {
	var row *models.User
	if err := r.DB.Where("id=?", user.ID).First(&row).Error; err != nil {
		return nil, err
	}

	err := r.DB.Model(&row).Select("*").Updates(&user).Error
	return row, err
}

func (r *userRepository) Delete(id string) error {
	var user *models.User
	if err := r.DB.Where("id=?", id).First(&user).Error; err != nil {
		return err
	}
	err := r.DB.Delete(&user).Error
	return err
}
