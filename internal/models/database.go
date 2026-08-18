package models

type DatabaseRequest struct {
	Name    string `json:"name,omitempty" example:"my_system"`
	OldName string `json:"oldName,omitempty" example:"db_teste"`
	NewName string `json:"newName,omitempty" example:"db_teste1"`
}

type DatabaseResponse struct {
	Active    string   `json:"active"`
	Databases []string `json:"databases"`
}