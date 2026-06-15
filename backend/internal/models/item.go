package models

type ItemRequest struct {
	Table  string         `json:"table"`
	Key    map[string]any `json:"key,omitempty"`
	Values map[string]any `json:"values,omitempty"`
}