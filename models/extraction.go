package models

type Extraction struct {
	ShipSymbol string          `json:"shipSymbol"`
	Yield      ExtractionYield `json:"yield"`
}
