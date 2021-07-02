package interfaces

import (
	"context"
	"github.com/and67o/go-comments/internal/configuration"
	"github.com/gorilla/mux"
	"net/http"
)

type HTTPApp interface {
	Start() error
	Stop(ctx context.Context) error
}

type Router interface {
	Hello(w http.ResponseWriter, _ *http.Request)
	CreateUser(w http.ResponseWriter, _ *http.Request)
	GetRouter() *mux.Router
}

type Config interface {
	GetHTTP() configuration.Server
}

type ClientError interface {
	Error() string
	ResponseBody() ([]byte, error)
	ResponseHeaders() (int, map[string]string)
}