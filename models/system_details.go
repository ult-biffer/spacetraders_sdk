package models

type SystemDetails struct {
	Coordinates
	Symbol       string     `json:"symbol"`
	SectorSymbol string     `json:"sectorSymbol"`
	Type         SystemType `json:"type"`
}
