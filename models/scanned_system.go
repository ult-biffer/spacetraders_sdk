package models

type ScannedSystem struct {
	SystemDetails
	Distance int `json:"distance"`
}
