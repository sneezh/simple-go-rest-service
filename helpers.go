package main

import (
	"fmt"
	"reflect"
)

func GetReflectedType(toReflect interface{}) (reflectedType reflect.Type) {
	reflectedType = reflect.TypeOf(toReflect).Elem()
	if reflectedType.Kind() == reflect.Slice {
		reflectedType = reflectedType.Elem()
	}
	return reflectedType
}

func printErrIfNotNil(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}
