package models

type ContractPayment struct {
	OnAccepted  int `json:"onAccepted"`
	OnFulfilled int `json:"onFulfilled"`
}
