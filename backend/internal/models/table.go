package models

type ColumnSchema struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	Nullable   bool   `json:"nullable"`
	Key        string `json:"key"`
	Default    string `json:"default"`
	Extra      string `json:"extra"`
	AutoNumber bool   `json:"autoNumber"`
}

type TableSchema struct {
	Name    string         `json:"name"`
	Columns []ColumnSchema `json:"columns"`
}

type TableData struct {
	Columns []string         `json:"columns"`
	Rows    []map[string]any `json:"rows"`
}