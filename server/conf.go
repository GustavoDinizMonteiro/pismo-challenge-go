package server

import (
	"github.com/gorilla/mux"
)

func CreateServer() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	return router
}
