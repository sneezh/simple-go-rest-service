package main

import (
	"encoding/json"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

func panicErrIfNotNil(statusCode *int, err error) {
	if err != nil {
		*statusCode = http.StatusInternalServerError
		printErrIfNotNil(err)
		panic(err)
	}
}

func notFoundIfErrNotNil(statusCode *int, err error) {
	if err != nil {
		*statusCode = http.StatusNotFound
		printErrIfNotNil(err)
	}
}

func getIdFromPath(pathVars map[string]string, varName string) int {
	Id, err := strconv.ParseInt(pathVars[varName], 10, 64)
	printErrIfNotNil(err)
	return int(Id)
}

func getBodyIntVariable(request *http.Request, varName string) int {
	variable, _ := strconv.ParseInt(request.URL.Query().Get(varName), 10, 64)
	return int(variable)
}

func getBodyStringVariable(request *http.Request, varName string) string {
	variable := request.URL.Query().Get(varName)
	return variable
}

func ReturnJson(writer http.ResponseWriter, response interface{}, statusCode *int) {
	if *statusCode > 299 || *statusCode == 201 || *statusCode == 204 {
		writer.WriteHeader(*statusCode)
		return
	}
	if reflect.TypeOf(response).Elem().Kind() == reflect.Slice && reflect.ValueOf(response).Elem().Len() < 1 {
		response = make([]int, 0)
	}
	if err := json.NewEncoder(writer).Encode(response); err != nil {
		panic(err)
	}
}

func handleErrorResponses(err error, statusCode *int) {
	if err != nil {
		switch err.Error() {
		case "field does not exists":
			*statusCode = http.StatusBadRequest
		default:
			*statusCode = http.StatusInternalServerError
		}
	}
}

func getDefaultStatus(method string) *int {
	statusCode := new(int)
	switch strings.ToUpper(method) {
	case "GET":
		*statusCode = http.StatusOK
	case "POST":
		*statusCode = http.StatusCreated
	case "PUT":
		*statusCode = http.StatusNoContent
	case "DELETE":
		*statusCode = http.StatusNoContent
	default:
		*statusCode = http.StatusNotImplemented
	}
	return statusCode
}
