package models

import "time"

type ContractTerms struct {
	Deadline time.Time `json:"deadline"`
	Payment  ContractPayment
	Deliver  []ContractDelivery
}
