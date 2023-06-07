package responses

import "spacetraders_sdk/models"

type CreateSurveyResponse struct {
	Data CreateSurveyResponseData `json:"data"`
}

type CreateSurveyResponseData struct {
	Cooldown models.Cooldown `json:"cooldown"`
	Surveys  []models.Survey `json:"survey"`
}
