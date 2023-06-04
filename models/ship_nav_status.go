package models

type ShipNavStatus string

const (
	STATUS_IN_TRANSIT ShipNavStatus = "IN_TRANSIT"
	STATUS_IN_ORBIT   ShipNavStatus = "IN_ORBIT"
	STATUS_DOCKED     ShipNavStatus = "DOCKED"
)
