package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"go-rest-webapp/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	uId, err := primitive.ObjectIDFromHex(p.ByName("id"))
	if err != nil {
		fmt.Println(err)
	}
	res := uc.client.Database("go-mongo").Collection("users").FindOne(r.Context(), bson.M{"_id": uId})

	u := models.User{}

	res.Decode(&u)

	uj, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)

}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}

	_ = json.NewDecoder(r.Body).Decode(&u)

	u.Id = primitive.NewObjectID()
	//ub, _ := bson.Marshal(u)

	res, err := uc.client.Database("go-mongo").Collection("users").InsertOne(r.Context(), bson.M{})
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User %s has been added!\n", res.InsertedID)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uId, err := primitive.ObjectIDFromHex(p.ByName("id"))
	if err != nil {
		fmt.Println(err)
	}
	res, err := uc.client.Database("go-mongo").Collection("users").DeleteOne(r.Context(), bson.M{"_id": uId})
	if err != nil {
		http.Error(w, "No user found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(res.DeletedCount), "users were deleted!")
}
