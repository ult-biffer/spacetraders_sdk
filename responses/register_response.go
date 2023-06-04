package responses

import "spacetraders_sdk/models"

type RegisterResponse struct {
	Data RegisterResponseData `json:"data"`
}

type RegisterResponseData struct {
	Agent    models.Agent    `json:"agent"`
	Contract models.Contract `json:"contract"`
	Faction  models.Faction  `json:"faction"`
	Ship     models.Ship     `json:"ship"`
	Token    string          `json:"token"`
}
