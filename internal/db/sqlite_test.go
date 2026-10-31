package db_test

import (
	"database/sql"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"gomyadm/internal/db"
	"gomyadm/internal/drivers"
	"gomyadm/internal/models"
	"gomyadm/internal/services"
)

func TestSQLiteConnectionListsTablesAndReadsData(t *testing.T) {
	path := filepath.Join(t.TempDir(), "sample database.sqlite")
	if err := os.WriteFile(path, nil, 0600); err != nil {
		t.Fatal(err)
	}

	driver := drivers.SQLiteDriver{}
	fixtureDB, err := sql.Open(driver.DriverName(), (&url.URL{
		Scheme:   "file",
		Path:     path,
		RawQuery: "mode=rwc",
	}).String())
	if err != nil {
		t.Fatal(err)
	}
	if _, err := fixtureDB.Exec(`
		CREATE TABLE users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			email TEXT UNIQUE
		);
		INSERT INTO users (name, email) VALUES ('Ada', 'ada@example.test');
	`); err != nil {
		fixtureDB.Close()
		t.Fatal(err)
	}
	if err := fixtureDB.Close(); err != nil {
		t.Fatal(err)
	}

	manager := db.NewConnectionManager()
	response, err := manager.Connect(models.ConnectionConfig{
		ID:       "sqlite-test",
		Name:     "SQLite test",
		Driver:   "sqlite",
		FilePath: path,
	})
	if err != nil {
		t.Fatalf("connect to SQLite file: %v", err)
	}
	defer manager.Disconnect()

	if _, err := manager.DB().Exec("INSERT INTO users (name) VALUES ('Grace')"); err == nil {
		t.Fatal("SQLite connection should be read-only")
	}

	if response.Database != "main" || !reflect.DeepEqual(response.Databases, []string{"main"}) {
		t.Fatalf("unexpected databases in connection response: %+v", response)
	}
	if response.FilePath != path {
		t.Fatalf("expected absolute file path %q, got %q", path, response.FilePath)
	}

	tables, err := services.NewTableService(manager).GetAll()
	if err != nil {
		t.Fatalf("list tables: %v", err)
	}
	if !reflect.DeepEqual(tables, []string{"users"}) {
		t.Fatalf("expected users table, got %v", tables)
	}

	data, err := services.NewItemService(manager).GetAll("users")
	if err != nil {
		t.Fatalf("read table rows: %v", err)
	}
	if !reflect.DeepEqual(data.Columns, []string{"id", "name", "email"}) {
		t.Fatalf("unexpected columns: %v", data.Columns)
	}
	if len(data.Rows) != 1 || data.Rows[0]["id"] != int64(1) || data.Rows[0]["name"] != "Ada" {
		t.Fatalf("unexpected table rows: %#v", data.Rows)
	}

	schema, err := services.NewColumnService(manager).GetAll("users")
	if err != nil {
		t.Fatalf("read table schema: %v", err)
	}
	if len(schema.Columns) != 3 ||
		schema.Columns[0].Name != "id" ||
		!schema.Columns[0].Primary ||
		!schema.Columns[0].AutoIncrement ||
		!schema.Columns[2].Unique {
		t.Fatalf("unexpected table schema: %+v", schema.Columns)
	}
}

func TestSQLiteConnectionRequiresExistingFile(t *testing.T) {
	manager := db.NewConnectionManager()
	_, err := manager.Connect(models.ConnectionConfig{
		Driver:   "sqlite",
		FilePath: filepath.Join(t.TempDir(), "missing.sqlite"),
	})
	if err == nil {
		t.Fatal("expected an error for a missing SQLite file")
	}

	if _, err := manager.Active(); err == nil {
		t.Fatal("connection should not be active after a failed attempt")
	}
}

func TestSQLiteConnectionRejectsNonDatabaseFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "not-a-database.sqlite")
	if err := os.WriteFile(path, []byte("not a SQLite database"), 0600); err != nil {
		t.Fatal(err)
	}

	manager := db.NewConnectionManager()
	_, err := manager.Connect(models.ConnectionConfig{
		Driver:   "sqlite",
		FilePath: path,
	})
	if err == nil {
		t.Fatal("expected an error for an invalid SQLite file")
	}
}
