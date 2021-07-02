package router

import (
	"errors"
	"github.com/and67o/go-comments/internal/app"
	"github.com/and67o/go-comments/internal/errors/che"
	"github.com/and67o/go-comments/internal/interfaces"
	"github.com/gorilla/mux"
	"io/ioutil"
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

func (r *Router) CreateUser(w http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		// может проще передать error
		che.Error(w,http.StatusInternalServerError, err.Error())
		return
	}
}

func New(app *app.App) interfaces.Router {
	var router Router

	router.app = app
	r := mux.NewRouter()
	router.initRoutes(r)

	router.muxRouter = r

	return &router
}

type appError struct {
	Error error
	Message string
	Code int
}

type appHandler func(http.ResponseWriter, *http.Request) *appError



func oleg(w http.ResponseWriter, r *http.Request) *appError  {
	return &appError{
		Error:   errors.New("oleg"),
		Message: "TEST",
		Code:    500,
	}
}

func (r *Router) initRoutes(router *mux.Router) {
	router.HandleFunc("/api/hello", r.Hello).Methods(http.MethodGet)

	router.HandleFunc("/api/register", r.CreateUser).Methods(http.MethodPost)
}
