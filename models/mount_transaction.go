package models

import "time"

type MountTransaction struct {
	TotalPrice int       `json:"totalPrice"`
	Timestamp  time.Time `json:"timestamp"`
}
