package user

import (
	"net/http"
	"rest-api/internal/handlers"

	"github.com/julienschmidt/httprouter"
)

const (
	usersURL = "/users"
	userURL = "/users/:uuid"
)


type handler struct {

}

func NewHandler() handlers.Handler {
	return &handler{}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(usersURL, h.GetList)
	router.GET(userURL, h.GetUserByUUID)
	router.GET(userURL, h.CreateUser)
	router.GET(userURL, h.UpdateUser)
	router.GET(userURL, h.PartiallyUpdateUser)
	router.GET(userURL, h.DeleteUser)
}


func (h *handler) GetList(w http.ResponseWriter, r * http.Request, params httprouter.Params) {
	w.Write([]byte("List of users"))
}

func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("Get user"))
}
func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("Create user"))
}
func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("Update user"))
}
func (h *handler) PartiallyUpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("Partially Update User"))
}
func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("Delete User"))
}





