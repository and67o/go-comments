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
	GetRouter() *mux.Router
}

type Config interface {
	GetHTTP() configuration.Server
}