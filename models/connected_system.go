package models

type ConnectedSystem struct {
	ScannedSystem
	FactionSymbol *string `json:"factionSymbol,omitempty"`
}
