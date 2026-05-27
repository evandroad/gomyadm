package models

type TableColumn struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	Nullable   bool   `json:"nullable"`
	PrimaryKey bool   `json:"primaryKey"`
}

type TableSchema struct {
	Name    string        `json:"name"`
	Columns []TableColumn `json:"columns"`
}