package models

type ScannedShipEngine struct {
	Symbol ShipEngineSymbol `json:"symbol"`
}

type ScannedShipFrame struct {
	Symbol ShipFrameSymbol `json:"symbol"`
}

type ScannedShipMount struct {
	Symbol ShipMountSymbol `json:"symbol"`
}

type ScannedShipReactor struct {
	Symbol ShipReactorSymbol `json:"symbol"`
}
