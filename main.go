package main

import (
	"context"
	"github.com/julienschmidt/httprouter"
	"go-rest-webapp/controllers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"time"
)

func main() {
	//docker run --name mongo -p 27017:27017 -e MONGO_INITDB_ROOT_USERNAME=root -e MONGO_INITDB_ROOT_PASSWORD=root -d mongo
	r := httprouter.New()
	uc := controllers.NewUserController(getClient())
	r.GET("/", index)
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	log.Fatal(http.ListenAndServe(":8000", r))
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	s := `<!DOCTYPE html>
			<html lang="en">
			<head>
			<meta charset="UTF-8">
			<title>Index</title>
			</head>
			<body>
			<a href="/user/9872309847">GO TO: http://localhost:8000/user/9872309847</a>
			</body>
			</html>`
	w.Header().Set("Content-Type", "text/html; charset=utf8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(s))
}

func getClient() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	c, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:root@localhost:27017"))
	if err != nil {
		panic(err)
	}
	return c
}
