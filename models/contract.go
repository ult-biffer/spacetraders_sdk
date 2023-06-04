package models

import "time"

type Contract struct {
	Id               string        `json:"id"`
	FactionSymbol    string        `json:"factionSymbol"`
	Type             string        `json:"type"`
	Terms            ContractTerms `json:"terms"`
	Accepted         bool          `json:"accepted"`
	Fulfilled        bool          `json:"fulfilled"`
	Expiration       time.Time     `json:"expiration"`
	DeadlineToAccept time.Time     `json:"deadlineToAccept"`
}
