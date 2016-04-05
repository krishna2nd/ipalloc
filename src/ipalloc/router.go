package ipalloc

import (
	"github.com/gorilla/mux"
)

func AppRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/search/{ip}", Search).Methods("GET")
	router.HandleFunc("/add", AddNewDevice).Methods("POST")
	return router
}
