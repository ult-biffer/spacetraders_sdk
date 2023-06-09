package util

import "strings"

type Location struct {
	Symbol   string
	Kind     LocationKind
	Sector   string
	System   string
	Waypoint string
}

type LocationKind int

const (
	LOC_INVALD LocationKind = iota
	LOC_SECTOR
	LOC_SYSTEM
	LOC_WAYPOINT
)

func NewLocation(symbol string) *Location {
	var (
		kind     LocationKind
		sector   string
		system   string
		waypoint string
	)

	parts := strings.Split(symbol, "-")

	switch len(parts) {
	case 1:
		kind = LOC_SECTOR
		sector = symbol
	case 2:
		kind = LOC_SYSTEM
		sector = parts[0]
		system = symbol
	case 3:
		kind = LOC_WAYPOINT
		sector = parts[0]
		system = strings.Join(parts[0:2], "-")
		waypoint = symbol
	}

	return &Location{
		Symbol:   symbol,
		Kind:     kind,
		Sector:   sector,
		System:   system,
		Waypoint: waypoint,
	}
}

func (loc *Location) IsWaypoint() bool {
	return loc.Kind == LOC_WAYPOINT
}

func (loc *Location) IsSystem() bool {
	return loc.Kind == LOC_SYSTEM
}

func (loc *Location) IsSector() bool {
	return loc.Kind == LOC_SECTOR
}

func (loc *Location) HasSystem() bool {
	return loc.System != ""
}
