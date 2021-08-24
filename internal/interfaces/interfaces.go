package interfaces

import (
	"context"
	"github.com/and67o/go-comments/internal/configuration"
	"github.com/and67o/go-comments/internal/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
	"time"
)

type HTTPApp interface {
	Start() error
	Stop(ctx context.Context) error
}

type Router interface {
	Hello(w http.ResponseWriter, _ *http.Request)
	CreateUser(w http.ResponseWriter, _ *http.Request)
	Login(w http.ResponseWriter, request *http.Request)
	LogOut(w http.ResponseWriter, request *http.Request)
	Refresh(w http.ResponseWriter, request *http.Request)
	GetRouter() *mux.Router
}

type Config interface {
	GetHTTP() configuration.Server
	GetDBConf() configuration.DBConf
	GetAuth() configuration.Auth
}

type Storage interface {
	Close() error
	GetDb() *gorm.DB
	UserService
}

type Redis interface {
	Set(key string, value interface{}, expiration time.Duration) error
	Get(key string) (interface{}, error)
}

type UserService interface {
	GetByEmail(email string) (*models.User, error)
	SaveUser(u models.User) (*models.User, error)
	GetById(id int64) (*models.User, error)
	GetUsers() (*[]models.User, error)
	DeleteUser(id uint64) error
}
