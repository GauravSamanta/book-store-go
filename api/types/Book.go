package types

import "gorm.io/datatypes"

type Book struct {
	Name        string
	Description string
	Tags        datatypes.JSONSlice[string]
	Author      string
}
