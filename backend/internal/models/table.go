package models

type TableSchema struct {
	Name    string             `json:"name"`
	Columns []ColumnDefinition `json:"columns"`
}

type TableData struct {
	Columns []string         `json:"columns"`
	Rows    []map[string]any `json:"rows"`
}