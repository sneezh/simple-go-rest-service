package main

import (
	. "github.com/gorilla/mux"
	"net/http"
)

const apiPrefix = ""

func getRouter() *Router {
	router := NewRouter()

	addRestRoutes(router, EntityApi{})

	return router
}

type RestApi interface {
	Path() string
	PathWithId() string
	One(writer http.ResponseWriter, request *http.Request)
	Many(writer http.ResponseWriter, request *http.Request)
	Create(writer http.ResponseWriter, request *http.Request)
	Update(writer http.ResponseWriter, request *http.Request)
	Delete(writer http.ResponseWriter, request *http.Request)
}

func addRestRoutes(router *Router, api RestApi) {
	router.HandleFunc(apiPrefix+api.Path(), api.Many).Methods("GET")
	router.HandleFunc(apiPrefix+api.Path(), api.Create).Methods("POST")
	router.HandleFunc(apiPrefix+api.PathWithId(), api.One).Methods("GET")
	router.HandleFunc(apiPrefix+api.PathWithId(), api.Update).Methods("PUT")
	router.HandleFunc(apiPrefix+api.PathWithId(), api.Delete).Methods("DELETE")
}
