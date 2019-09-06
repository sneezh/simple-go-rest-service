package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

const (
	EntityPath       = "/entities"
	EntityPathWithId = "/entities/{entityId}"
)

type EntityApi struct{}

func (api EntityApi) Path() string {
	return EntityPath
}

func (api EntityApi) PathWithId() string {
	return EntityPathWithId
}

func (api EntityApi) One(writer http.ResponseWriter, request *http.Request) {
	statusCode := getDefaultStatus(request.Method)
	entityToResponse := new(EntityResponse)
	defer ReturnJson(writer, entityToResponse, statusCode)

	pathVars := mux.Vars(request)
	entityToResponse.Id = getIdFromPath(pathVars, "entityId")

	err := entityToResponse.Get()
	notFoundIfErrNotNil(statusCode, err)
}

func (api EntityApi) Many(writer http.ResponseWriter, request *http.Request) {
	statusCode := getDefaultStatus(request.Method)
	response := new(EntitiesResponse)
	defer ReturnJson(writer, response, statusCode)

	limit := getBodyIntVariable(request, "limit")
	offset := getBodyIntVariable(request, "offset")

	err := getAllEntities(limit, offset, response)

	printErrIfNotNil(err)

	handleErrorResponses(err, statusCode)
}

func (api EntityApi) Create(writer http.ResponseWriter, request *http.Request) {
	statusCode := getDefaultStatus(request.Method)
	defer ReturnJson(writer, nil, statusCode)

	entity := new(Entity)
	err := json.NewDecoder(request.Body).Decode(&entity)
	panicErrIfNotNil(statusCode, err)

	err = entity.Create()
	panicErrIfNotNil(statusCode, err)
}

func (api EntityApi) Update(writer http.ResponseWriter, request *http.Request) {
	statusCode := getDefaultStatus(request.Method)
	defer ReturnJson(writer, nil, statusCode)

	entity := new(Entity)
	err := json.NewDecoder(request.Body).Decode(&entity)
	panicErrIfNotNil(statusCode, err)

	pathVars := mux.Vars(request)
	entityId := getIdFromPath(pathVars, "entityId")

	if entity.Id != entityId {
		*statusCode = http.StatusBadRequest
		return
	}

	err = entity.Update()
	panicErrIfNotNil(statusCode, err)
}

func (api EntityApi) Delete(writer http.ResponseWriter, request *http.Request) {
	statusCode := getDefaultStatus(request.Method)
	defer ReturnJson(writer, nil, statusCode)

	pathVars := mux.Vars(request)
	entityId := getIdFromPath(pathVars, "entityId")

	entity := new(Entity)
	entity.Id = entityId
	err := entity.Get()
	if err != nil {
		*statusCode = http.StatusNotFound
		return
	}

	err = entity.Delete()
	panicErrIfNotNil(statusCode, err)

	*statusCode = http.StatusNoContent
}
