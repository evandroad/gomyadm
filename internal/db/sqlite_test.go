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

	if _, err := manager.DB().Exec("INSERT INTO users (name) VALUES ('Grace')"); err != nil {
		t.Fatalf("SQLite connection should allow writes: %v", err)
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
	if len(data.Rows) != 2 || data.Rows[0]["id"] != int64(1) || data.Rows[0]["name"] != "Ada" {
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

func TestSQLiteWriteOperations(t *testing.T) {
	path := filepath.Join(t.TempDir(), "writable.sqlite")
	if err := os.WriteFile(path, nil, 0600); err != nil {
		t.Fatal(err)
	}

	manager := db.NewConnectionManager()
	_, err := manager.Connect(models.ConnectionConfig{Driver: "sqlite", FilePath: path})
	if err != nil {
		t.Fatalf("connect to SQLite file: %v", err)
	}
	defer manager.Disconnect()

	driver, ok := drivers.GetDriver("sqlite")
	if !ok {
		t.Fatal("SQLite driver was not registered")
	}

	if err := driver.CreateTable(manager.DB(), models.Table{
		Name: "people",
		Columns: []models.Column{
			{Name: "id", Type: "INTEGER", Primary: true, AutoIncrement: true},
			{Name: "name", Type: "TEXT", Nullable: false},
		},
	}); err != nil {
		t.Fatalf("create table: %v", err)
	}

	if err := driver.CreateItem(manager.DB(), "people", map[string]any{"name": "Ada"}); err != nil {
		t.Fatalf("insert row: %v", err)
	}
	item, err := driver.GetOneItem(manager.DB(), "people", map[string]any{"id": int64(1)})
	if err != nil || item["name"] != "Ada" {
		t.Fatalf("read inserted row: item=%v err=%v", item, err)
	}

	if err := driver.UpdateItem(manager.DB(), "people", map[string]any{"id": int64(1)}, map[string]any{"name": "Grace"}); err != nil {
		t.Fatalf("update row: %v", err)
	}
	item, err = driver.GetOneItem(manager.DB(), "people", map[string]any{"id": int64(1)})
	if err != nil || item["name"] != "Grace" {
		t.Fatalf("read updated row: item=%v err=%v", item, err)
	}

	if err := driver.CreateColumn(manager.DB(), "people", models.Column{
		Name: "email", Type: "TEXT", Nullable: true, Unique: true,
	}); err != nil {
		t.Fatalf("add column: %v", err)
	}
	if err := driver.UpdateColumn(manager.DB(), "people", "email", models.Column{
		Name: "contact", Type: "TEXT", Nullable: true, Unique: true,
	}); err != nil {
		t.Fatalf("rename column: %v", err)
	}
	if err := driver.DeleteColumn(manager.DB(), "people", "contact"); err != nil {
		t.Fatalf("drop column: %v", err)
	}

	if err := driver.UpdateTable(manager.DB(), "people", "users"); err != nil {
		t.Fatalf("rename table: %v", err)
	}
	if err := driver.DeleteItem(manager.DB(), "users", map[string]any{"id": int64(1)}); err != nil {
		t.Fatalf("delete row: %v", err)
	}
	if err := driver.DeleteTable(manager.DB(), "users"); err != nil {
		t.Fatalf("drop table: %v", err)
	}

	if err := driver.CreateTable(manager.DB(), models.Table{
		Name: "notes",
		Columns: []models.Column{
			{Name: "body", Type: "TEXT", Nullable: false},
		},
	}); err != nil {
		t.Fatalf("create table without primary key: %v", err)
	}
	if err := driver.CreateItem(manager.DB(), "notes", map[string]any{"body": "draft"}); err != nil {
		t.Fatalf("insert row without primary key: %v", err)
	}
	notes, err := driver.GetAllItem(manager.DB(), "notes")
	if err != nil {
		t.Fatalf("read row without primary key: %v", err)
	}
	if len(notes.Rows) != 1 || notes.Rows[0]["__gomyadm_rowid"] == nil {
		t.Fatalf("expected hidden rowid for table without primary key, got %#v", notes.Rows)
	}
	rowID := notes.Rows[0]["__gomyadm_rowid"]
	if err := driver.UpdateItem(manager.DB(), "notes", map[string]any{"__gomyadm_rowid": rowID}, map[string]any{"body": "published"}); err != nil {
		t.Fatalf("update row without primary key: %v", err)
	}
	if err := driver.DeleteItem(manager.DB(), "notes", map[string]any{"__gomyadm_rowid": rowID}); err != nil {
		t.Fatalf("delete row without primary key: %v", err)
	}
	if err := driver.DeleteTable(manager.DB(), "notes"); err != nil {
		t.Fatalf("drop table without primary key: %v", err)
	}

	tables, err := driver.GetAllTable(manager.DB())
	if err != nil || len(tables) != 0 {
		t.Fatalf("expected no remaining tables, got %v (err=%v)", tables, err)
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
