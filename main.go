package main

import (
	"github.com/julienschmidt/httprouter"
	"go-rest-webapp/controllers"
	"log"
	"net/http"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController()
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
