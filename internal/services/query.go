package services

import (
	"database/sql"
	"fmt"
	"strings"

	"gomyadm/internal/db"
	"gomyadm/internal/models"
)

type QueryService struct {
	Connection *db.ConnectionManager
}

func NewQueryService(conn *db.ConnectionManager) *QueryService {
	return &QueryService{
		Connection: conn,
	}
}

func (s *QueryService) ExecuteQuery(query string) (*models.QueryResult, error) {
	if s.Connection.GetDatabase() == "" {
		return nil, fmt.Errorf("No database selected.")
	}

	if isQuery(query) {
		return executeSelect(s.Connection.DB(), query)
	}

	return executeCommand(s.Connection.DB(), query)
}

func isQuery(sql string) bool {
	sql = strings.TrimSpace(strings.ToUpper(sql))

	return strings.HasPrefix(sql, "SELECT") ||
		strings.HasPrefix(sql, "SHOW") ||
		strings.HasPrefix(sql, "DESCRIBE") ||
		strings.HasPrefix(sql, "EXPLAIN") ||
		strings.HasPrefix(sql, "WITH")
}

func executeSelect(db *sql.DB, query string) (*models.QueryResult, error) {
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	result := &models.QueryResult{
		Type:    "query",
		Columns: columns,
		Rows:    []map[string]interface{}{},
	}

	for rows.Next() {
		values := make([]any, len(columns))
		pointers := make([]any, len(columns))

		for i := range values {
			pointers[i] = &values[i]
		}

		if err := rows.Scan(pointers...); err != nil {
			return nil, err
		}

		row := make(map[string]interface{})

		for i, col := range columns {
			v := values[i]

			if b, ok := v.([]byte); ok {
				row[col] = string(b)
			} else {
				row[col] = v
			}
		}

		result.Rows = append(result.Rows, row)
	}

	return result, nil
}

func executeCommand(db *sql.DB, query string) (*models.QueryResult, error) {
	res, err := db.Exec(query)
	if err != nil {
		return nil, err
	}

	affected, _ := res.RowsAffected()

	return &models.QueryResult{
		Type:         "command",
		RowsAffected: affected,
	}, nil
}