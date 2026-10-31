package bindings

import (
	"context"
	"fmt"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type FileDialogBinding struct {
	ctx context.Context
}

func NewFileDialog() *FileDialogBinding {
	return &FileDialogBinding{}
}

func (b *FileDialogBinding) SetContext(ctx context.Context) {
	b.ctx = ctx
}

func (b *FileDialogBinding) OpenSQLiteFile() (string, error) {
	if b.ctx == nil {
		return "", fmt.Errorf("file dialog is not available before application startup")
	}

	return runtime.OpenFileDialog(b.ctx, runtime.OpenDialogOptions{
		Title: "Selecionar arquivo SQLite",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Arquivos SQLite (*.db, *.sqlite, *.sqlite3, *.db3)",
				Pattern:     "*.db;*.sqlite;*.sqlite3;*.db3",
			},
		},
	})
}
