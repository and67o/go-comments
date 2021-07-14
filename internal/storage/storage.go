package storage

import (
	"fmt"
	"github.com/and67o/go-comments/internal/configuration"
	"github.com/and67o/go-comments/internal/interfaces"
	"github.com/and67o/go-comments/internal/service"
	_ "github.com/go-sql-driver/mysql" // nolint: gci
	"github.com/jinzhu/gorm"
)

type Storage struct {
	db *gorm.DB
	service *interfaces.UserService
}

const driverName = "mysql"
const format = "2006-01-02 15:04:05"

func New(config configuration.DBConf) (interfaces.Storage, error) {
	db, err := gorm.Open(driverName, dataSourceName(config))
	if err != nil {
		return nil, fmt.Errorf("connect db: %w", err)
	}
	return &Storage{
		db: db,
		service: service.NewUserService(db),
	}, nil
}

func (s *Storage) Close() error {
	err := s.db.Close()
	if err != nil {
		return fmt.Errorf("close connect: %w", err)
	}

	return nil
}

func (s *Storage) GetDb() *gorm.DB {
	return s.db
}

func (s *Storage) GetService() *service.UserService {
	return s.service
}

func dataSourceName(config configuration.DBConf) string {
	return fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)
}

