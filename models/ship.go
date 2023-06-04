package models

type Ship struct {
	Symbol       string           `json:"symbol"`
	Registration ShipRegistration `json:"registration"`
	Nav          ShipNav          `json:"nav"`
	Crew         ShipCrew         `json:"crew"`
	Frame        ShipFrame        `json:"frame"`
	Reactor      ShipReactor      `json:"reactor"`
	Engine       ShipEngine       `json:"engine"`
	Modules      []ShipModule     `json:"modules"`
	Mounts       []ShipMount      `json:"mounts"`
	Cargo        ShipCargo        `json:"cargo"`
	Fuel         ShipFuel         `json:"fuel"`
}
