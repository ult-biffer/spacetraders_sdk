package models

type JumpGate struct {
	JumpRange        int               `json:"jumpRange"`
	FactionSymbol    *string           `json:"factionSymbol,omitempty"`
	ConnectedSystems []ConnectedSystem `json:"connectedSystems"`
}
