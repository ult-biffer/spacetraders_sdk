package models

type Describable[T any] struct {
	Symbol      T      `json:"symbol"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
