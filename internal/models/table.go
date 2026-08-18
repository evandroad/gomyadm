package models

type Table struct {
	Name    string   `json:"name" example:"teste"`
	Columns []Column `json:"columns"`
}

type TableData struct {
	Columns []string         `json:"columns"`
	Rows    []map[string]any `json:"rows"`
}