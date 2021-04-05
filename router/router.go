package router

import (
	"github.com/agniswarm/go-lambda/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/book", controller.Show).Methods("GET")
	r.HandleFunc("/book", controller.CreateBook).Methods("POST")
	return r
}
