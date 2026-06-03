package models

type QueryResult struct {
	Type         string                   `json:"type"`
	Columns      []string                 `json:"columns,omitempty"`
	Rows         []map[string]interface{} `json:"rows,omitempty"`
	RowsAffected int64                    `json:"rowsAffected,omitempty"`
}

type QueryRequest struct {
	Query string `json:"query"`
}