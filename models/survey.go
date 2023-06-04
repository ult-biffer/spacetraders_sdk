package models

import "time"

type Survey struct {
	Signature  string          `json:"signature"`
	Symbol     string          `json:"symbol"`
	Deposits   []SurveyDeposit `json:"deposits"`
	Expiration time.Time       `json:"expiration"`
	Size       SurveySize      `json:"size"`
}
