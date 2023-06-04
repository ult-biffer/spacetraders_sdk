package models

type Faction struct {
	Describable[string]
	Headquarters string         `json:"headquarters"`
	Traits       []FactionTrait `json:"traits"`
	IsRecruiting bool           `json:"isRecruiting"`
}
