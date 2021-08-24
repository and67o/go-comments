package storage

import (
	"github.com/and67o/go-comments/internal/models"
)

func (s *Storage) GetByEmail(email string) (*models.User, error) {
	user := models.User{}

	err:= s.db.Where(&models.User{
		Login:    email,
	}).First(&user).Error
	if err != nil {
		return &user, err
	}

	return &user, err
}

func (s *Storage) SaveUser(u models.User) (*models.User, error) {
	var err error

	err = s.db.
		Model(models.User{}).
		Create(&u).Error
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (s *Storage) GetById(id int64) (*models.User, error) {
	var user models.User

	err := s.db.First(&user, id).Error
	if err != nil {
		return &user, err
	}

	return &user, err
}

func (s *Storage) GetUsers() (*[]models.User, error) {
	var users []models.User
	err := s.db.Find(&users).Error
	if err != nil {
		return &users, err
	}

	return &users, err
}

func (s *Storage) DeleteUser(id uint64) error {
	err := s.db.Delete(models.User{}, id).Error
	if err != nil {
		return err
	}

	return nil
}
