package main

import (
	"log"
	"net/http"
	"time"

	data_user "github.com/claudiomozer/muxgrpc/internal/data/user"
	"github.com/claudiomozer/muxgrpc/internal/presentation/controllers"
	"github.com/claudiomozer/muxgrpc/pkg/infra/repositories/static"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	userRepository := static.NewUserRepository()
	service := data_user.New(userRepository, userRepository)
	createUserController := controllers.NewCreateUserController(service)
	findUserByDocument := controllers.NewFindUserByDocumentController(service)

	r.Handle("/users", createUserController).Methods("POST")
	r.Handle("/users/{document}", findUserByDocument).Methods("GET")
	http.Handle("/", r)

	server := getHTTPServer(r)
	log.Fatal(server.ListenAndServe())
}

func getHTTPServer(router *mux.Router) *http.Server {
	return &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: time.Minute,
		ReadTimeout:  time.Minute,
	}
}
