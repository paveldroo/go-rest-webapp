package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"go-rest-webapp/models"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type UserController struct {
	client *mongo.Client
}

func NewUserController(c *mongo.Client) *UserController {
	return &UserController{c}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{
		Name:   "Jonny",
		Gender: "male",
		Age:    28,
		Id:     p.ByName("id"),
	}

	uj, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)

}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}

	_ = json.NewDecoder(r.Body).Decode(&u)

	u.Id = "007"
	uj, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// We will delete user here
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User was deleted!")
}
