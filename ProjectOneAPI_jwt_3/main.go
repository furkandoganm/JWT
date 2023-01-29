package main

import (
	"net/http"
	"projects/ProjectOneAPI_jwt_3/service"
	"projects/ProjectOneAPI_jwt_3/service/collection"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/login", service.Login)

	mux.HandleFunc("/getmovies", service.IsAuthorized(collection.GetMovie))
	mux.HandleFunc("/addmovie", service.IsAuthorized(collection.AddMovie))

	http.ListenAndServe(":8080", mux)
}
