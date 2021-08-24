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

func (r *Router) LogOut(w http.ResponseWriter, request *http.Request) {
	panic("implement me")
}

func (r *Router) Refresh(w http.ResponseWriter, request *http.Request) {
	panic("implement me")
}

func (r *Router) Hello(w http.ResponseWriter, _ *http.Request) {
	panic("implement me")
}

func (r *Router) GetRouter() *mux.Router {
	return r.muxRouter
}

func (r *Router) Login(w http.ResponseWriter, request *http.Request) {
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

	userDb, err := r.app.Storage.GetByEmail(user.Login)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, errors.New("email busy"))
		return
	}

	err = token.VerifyPassword(user.Password, userDb.Password)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, err)
	}

	tokens, err := token.GetTokens(userDb.Id, r.app.Config.GetAuth())
	if err != nil {
		return
	}

	err = r.app.Redis.Set(
		token.GetAccessKey(userDb.Id),
		userDb.Id,
		time.Unix(int64(r.app.Config.GetAuth().AccessExpire), 0).Sub(time.Now()),
	)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
	}

	err = r.app.Redis.Set(
		token.GetRefreshKey(userDb.Id),
		userDb.Id,
		time.Unix(int64(r.app.Config.GetAuth().RefreshExpire), 0).Sub(time.Now()),
	)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
	}

	response.Json(w, http.StatusOK, tokens)
}

func (r *Router) CreateUser(w http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	defer request.Body.Close()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	if len(body) == 0 {
		response.Error(w, http.StatusBadRequest, errors.New("empty request"))
		return
	}

	var user models.User

	err = json.Unmarshal(body, &user)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	_, err = r.app.Storage.GetByEmail(user.Login)
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

	newUser, err := r.app.Storage.SaveUser(user)
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

func (r *Router) initRoutes(router *mux.Router) {
	router.HandleFunc("/api/hello", r.Hello).Methods(http.MethodGet)

	router.HandleFunc("/api/register", r.CreateUser).Methods(http.MethodPost)
}
