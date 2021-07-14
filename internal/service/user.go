package service

import (
	"github.com/and67o/go-comments/internal/interfaces"
	"github.com/and67o/go-comments/internal/models"
	"github.com/jinzhu/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) interfaces.UserService {
	var service UserService

	service.db = db
	return &service
}

func (s *UserService) GetByEmail(email string) (*models.User, error) {
	user := models.User{}

	err := s.db.Model(models.User{}).
		Where("email = ?", email).
		Take(&user).
		Error

	if err != nil {
		return &user, err
	}

	return &user, err
}

func (s *UserService) SaveUser(u models.User) (*models.User, error) {
	var err error

	err = s.db.Create(&u).Error
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (s *UserService) GetById(id int64) (*models.User, error) {
	var user models.User

	err := s.db.First(&user, id).Error
	if err != nil {
		return &user, err
	}

	return &user, err
}

func (s *UserService) GetUsers() (*[]models.User, error) {
	var users []models.User
	err := s.db.Find(&users).Error
	if err != nil {
		return &users, err
	}

	return &users, err
}

func (s *UserService) DeleteUser(id uint64) error {
	err := s.db.Delete(models.User{}, id).Error
	if err != nil {
		return err
	}

	return nil
}
