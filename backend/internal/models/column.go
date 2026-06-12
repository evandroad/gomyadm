package models

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