package main

import "time"

type InfoResponse struct {
	ServiceName    string    `json:"product_name"`
	ServiceVersion string    `json:"version"`
	ReleaseDate    time.Time `json:"updated_at"`
}

type EntityResponse struct {
	Entity
}

type EntitiesResponse []Entity
