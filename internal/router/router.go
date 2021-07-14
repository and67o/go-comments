package router

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/and67o/go-comments/internal/app"
	"github.com/and67o/go-comments/internal/errors/response"
	"github.com/and67o/go-comments/internal/hash"
	"github.com/and67o/go-comments/internal/interfaces"
	"github.com/and67o/go-comments/internal/models"
	"github.com/and67o/go-comments/internal/token"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"time"
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
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	var user models.User

	err = json.Unmarshal(body, &user)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
	}

	_, err = r.app.Storage.GetService().GetByEmail(user.Login)
	if err == nil {
		response.Error(w, http.StatusConflict, errors.New("email busy"))
		return
	}

	hashedPassword, err := hash.HashPassword(user.Password)
	if err != nil {
		response.Error(w, http.StatusBadRequest, errors.New("hash error"))
		return
	}
	user.Password = hashedPassword

	newUser, err := r.app.Storage.GetService().SaveUser(user)
	if err != nil {
		response.Error(w, http.StatusBadRequest, errors.New(fmt.Sprintf("create user err: %v", err.Error())))
		return
	}

	tokens, err := token.GetTokens(newUser.Id, r.app.Config.GetAuth())
	if err != nil {
		return
	}

	err = r.app.Redis.Set(
		token.GetAccessKey(newUser.Id),
		newUser.Id,
		time.Unix(int64(r.app.Config.GetAuth().AccessExpire), 0).Sub(time.Now()),
	)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
	}

	err = r.app.Redis.Set(
		token.GetRefreshKey(newUser.Id),
		newUser.Id,
		time.Unix(int64(r.app.Config.GetAuth().RefreshExpire), 0).Sub(time.Now()),
	)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
	}

	response.Json(w, http.StatusOK, tokens)
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
	Error   error
	Message string
	Code    int
}

type appHandler func(http.ResponseWriter, *http.Request) *appError

func oleg(w http.ResponseWriter, r *http.Request) *appError {
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
