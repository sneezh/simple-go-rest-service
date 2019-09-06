package main

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func init() {
	isTesting = true
}

func TestGetReflectedType(t *testing.T) {
	type Example struct {
	}
	type Examples []Example

	singleStrict := new(Example)
	reflectedType := GetReflectedType(singleStrict)
	expectedType := reflect.TypeOf(singleStrict).Elem()
	assert.Equal(t, expectedType, reflectedType)

	sliceStruct := new(Examples)
	reflectedType = GetReflectedType(sliceStruct)
	expectedType = reflect.TypeOf(sliceStruct).Elem().Elem()
	assert.Equal(t, expectedType, reflectedType)
}
