package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"go-rest-webapp/models"
	"log"
	"net/http"
)

func main() {
	r := httprouter.New()
	r.GET("/", index)
	r.GET("/user/:id", getUser)
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

func getUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{
		Name:   "Jonny",
		Gender: "male",
		Age:    28,
		Id:     p.ByName("id"),
	}

	uj, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s\n", uj)

}
