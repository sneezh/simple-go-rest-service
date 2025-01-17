package main

import "time"

type Timestamps struct {
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `pg:",soft_delete" json:"deleted_at,omitempty"`
}
