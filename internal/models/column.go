package models

import (
	"encoding/json"
)

type Column struct {
	Name          string `json:"name" example:"teste"`
	Type          string `json:"type" example:"INT"`
	Length        *int   `json:"length,omitempty" example:"10"`
	Nullable      bool   `json:"nullable" example:"true"`
	Primary       bool   `json:"primary" example:"false"`
	Unique        bool   `json:"unique" example:"true"`
	AutoIncrement bool   `json:"autoIncrement" example:"false"`
	DefaultValue  string `json:"defaultValue" example:""`
}

func (c *Column) ToString() string {
	v, _ := json.Marshal(c)
	return string(v)
}

type ColumnRequest struct {
	Table   string `json:"table"`
	OldName string `json:"oldName"`
	Column  Column `json:"column"`
}