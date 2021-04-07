package router

import (
	"net/http"

	"github.com/agniswarm/go-lambda/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	bookRouter := r.PathPrefix("/book").Subrouter()
	bookRouter.HandleFunc("/read", controller.Show).Methods(http.MethodGet)
	bookRouter.HandleFunc("/create", controller.CreateBook).Methods(http.MethodPost)
	return r
}
