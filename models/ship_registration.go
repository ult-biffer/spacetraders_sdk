package models

type ShipRegistration struct {
	Name          string   `json:"name"`
	FactionSymbol string   `json:"factionSymbol"`
	Role          ShipRole `json:"role"`
}
