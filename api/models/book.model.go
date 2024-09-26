package models

import "gorm.io/datatypes"

type Book struct {
	Id          uint                        `gorm:"primaryKey"`
	Name        string                      `json:"name"`
	Description string                      `json:"description"`
	Tags        datatypes.JSONSlice[string] `json:"tags"`
	Author      string                      `json:"author"`
}
