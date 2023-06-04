package models

type Shipyard struct {
	Symbol       string                `json:"symbol"`
	ShipTypes    []ShipType            `json:"shipTypes"`
	Transactions []ShipyardTransaction `json:"transactions,omitempty"`
	Ships        []ShipyardShip        `json:"ships,omitempty"`
}
