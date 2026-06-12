package models

import (
	"encoding/json"
)

type ColumnDefinition struct {
	Name          string `json:"name"`
	Type          string `json:"type"`
	Length        *int   `json:"length,omitempty"`
	Nullable      bool   `json:"nullable"`
	Primary       bool   `json:"primary"`
	Unique        bool   `json:"unique"`
	AutoIncrement bool   `json:"autoIncrement"`
	DefaultValue  string `json:"defaultValue"`
}

func (c *ColumnDefinition) ToString() string {
	v, _ := json.Marshal(c)
	return string(v)
}