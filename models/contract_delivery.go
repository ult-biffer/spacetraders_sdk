package models

type ContractDelivery struct {
	TradeSymbol       string `json:"tradeSymbol"`
	DestinationSymbol string `json:"destinationSymbol"`
	UnitsRequired     int    `json:"unitsRequired"`
	UnitsFulfilled    int    `json:"unitsFulfilled"`
}
