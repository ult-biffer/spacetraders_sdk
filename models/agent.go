package models

type Agent struct {
	AccountId       string `json:"accountId"`
	Symbol          string `json:"symbol"`
	Headquarters    string `json:"headquarters"`
	Credits         int    `json:"credits"`
	StartingFaction string `json:"startingFaction"`
}
