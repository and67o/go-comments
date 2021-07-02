package router

import (
	"github.com/and67o/go-comments/internal/app"
	"github.com/and67o/go-comments/internal/interfaces"
	"github.com/gorilla/mux"
	"net/http"
)

type Router struct {
	muxRouter *mux.Router
	app       *app.App
}

func (r *Router) Hello(w http.ResponseWriter, _ *http.Request) {
	panic("implement me")
}

func (r *Router) GetRouter() *mux.Router {
	return r.muxRouter
}

func New(app *app.App) interfaces.Router {
	var router Router

	router.app = app
	r := mux.NewRouter()
	router.initRoutes(r)

	router.muxRouter = r

	return &router
}

func (r *Router) initRoutes(router *mux.Router) {

}
